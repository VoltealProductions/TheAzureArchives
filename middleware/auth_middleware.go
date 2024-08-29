package middleware

import (
	"net/http"

	"github.com/VoltealProductions/TheAzureArcchives/services/auth"
	"github.com/VoltealProductions/TheAzureArcchives/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		_, err := auth.VerifyToken(token)
		if err != nil {
			utils.WriteError(w, http.StatusUnauthorized, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}
