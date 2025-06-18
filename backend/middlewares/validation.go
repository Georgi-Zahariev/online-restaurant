package middlewares

import (
	"encoding/json"
	"net/http"
)

func ValidationMiddleware(expectedFields []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var input map[string]interface{}
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
				http.Error(w, "Invalid JSON input", http.StatusBadRequest)
				return
			}

			for _, field := range expectedFields {
				if _, exists := input[field]; !exists {
					http.Error(w, "Missing required field: "+field, http.StatusBadRequest)
					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}
