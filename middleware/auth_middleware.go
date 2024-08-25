package middleware

import (
	"net/http"
)

func AuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "Bearer token" {
			http.Error(w, "Unathorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
