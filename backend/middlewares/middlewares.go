package middlewares

import (
	"encoding/json"
	"net/http"
)

// Middleware to set "Content-Type" to "application/json"
func JSONContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// Middleware to validate "Authorization" header with a hardcoded token
func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader != "Bearer abc" {
			w.WriteHeader(http.StatusUnauthorized)
			response := map[string]string{"error": "Unauthorized: Invalid or missing Authorization header"}
			json.NewEncoder(w).Encode(response)
			return
		}
		next.ServeHTTP(w, r)
	})
}
