package web

import (
	"encoding/json"
	"gateway/internals/domain"
	"gateway/internals/http/entities"
	"gateway/internals/services"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type InvoiceHandler struct {
	invoiceService *services.InvoiceService
}

func NewInvoiceHandler(invoiceService *services.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{
		invoiceService: invoiceService,
	}
}

// Endpoint: /invoice
// Method: POST
// Requer autenticação via X-API-KEY

func (h *InvoiceHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input entities.CreateInvoiceInput
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	input.APIKey = r.Header.Get("X-API-KEY")



	output, err := h.invoiceService.Create(input)


	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)


}


func (h *InvoiceHandler) Get(w http.ResponseWriter, r *http.Request) {
	id :=  chi.URLParam(r, "id")

	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	apiKey := r.Header.Get("X-API-KEY")

	if apiKey == "" {
		http.Error(w, "API KEY is required", http.StatusUnauthorized)
		return
	}

	output, err := h.invoiceService.GetByID(id, apiKey)

	if err != nil {
		switch err {
		case domain.ErrInvoiceNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		case domain.ErrAccountNotFound:
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		case domain.ErrUnauthorized:
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		default:	
			http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

// ListByAccount

func (h *InvoiceHandler) ListByAccount(w http.ResponseWriter, r *http.Request) {

	apiKey := r.Header.Get("X-API-KEY")
	if apiKey == "" {
		http.Error(w, "API KEY is required", http.StatusUnauthorized)
		return
	}

	accountID := chi.URLParam(r, "accountID")

	if accountID == "" {
		http.Error(w, "Account ID is required", http.StatusBadRequest)
		return
	}

	output, err := h.invoiceService.ListByAccount(accountID)

	if err != nil {
		switch err {
		case domain.ErrAccountNotFound:
			http.Error(w, err.Error(), http.StatusUnauthorized)
			default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)

}}