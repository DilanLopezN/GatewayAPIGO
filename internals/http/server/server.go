package server

import (
	"gateway/internals/http/middlewares"
	"gateway/internals/http/web"
	"gateway/internals/services"
	"net/http"

	"github.com/go-chi/chi/v5"
)


type Server struct {
	router *chi.Mux
	server *http.Server
	accountService *services.AccountService
	invoiceService *services.InvoiceService
	port string
}

func NewServer(accountService *services.AccountService, 
	invoiceService *services.InvoiceService,
	port string) *Server {
	return &Server{
		router: chi.NewRouter(),
		accountService: accountService,
		invoiceService: invoiceService,
		port: port,
}}


func (s *Server) ConfigureRoutes() {
	accountHandler := web.NewAccountHandler(s.accountService)
	invoiceHandler := web.NewInvoiceHandler(s.invoiceService)
	authMiddleware := middlewares.NewAuthMiddleware(s.accountService)


	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is running"))
	})

	s.router.Post("/accounts", accountHandler.Create)
	s.router.Get("/accounts", accountHandler.Get)

	s.router.Group(func(r chi.Router) {
		r.Use(authMiddleware.Authenticate)
		s.router.Post("/invoice", invoiceHandler.Create)
		s.router.Get("/invoice", invoiceHandler.ListByAccount)
		s.router.Get("/invoice/{id}", invoiceHandler.Get)
	
	} )



}


func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}

	return s.server.ListenAndServe()
}