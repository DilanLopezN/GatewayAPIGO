package web

import (
	"encoding/json"
	"gateway/internals/http/entities"
	"gateway/internals/services"
	"net/http"
)

type AccountHandler struct {
	accountService *services.AccountService

}

func NewAccountHandler(accountService *services.AccountService) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
	}
}

func (h *AccountHandler) Create(w http.ResponseWriter, r *http.Request) {
	
	var input entities.CreateAccountInput

	if err := json.NewDecoder(r.Body).Decode(&input); 
	err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account, err := h.accountService.CreateAccount(input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)

}