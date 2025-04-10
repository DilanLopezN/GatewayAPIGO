package middlewares

import (
	"gateway/internals/services"
	"net/http"
)

type AuthMiddleware struct {
	accountService services.AccountService
}

func NewAuthMiddleware(accountService services.AccountService) *AuthMiddleware {
	return &AuthMiddleware{
		accountService: accountService,
	}
}


// AUTHENTICATE
func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")

		if apiKey == "" {
			http.Error(w, "API KEY is required", http.StatusUnauthorized)
			return
		}

		account, err := m.accountService.FindByAPIKey(apiKey)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if account == nil {
			http.Error(w, "API KEY not found", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})


}
