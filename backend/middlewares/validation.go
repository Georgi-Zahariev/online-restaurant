package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Georgi-Zahariev/online-restaurant/backend/models"
)

func ValidationMiddleware(entity string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var instance interface {
				Validate() error
			}

			switch entity {
			case "users":
				instance = &models.User{}
			case "dishes":
				instance = &models.Dish{}
			case "orders":
				instance = &models.Order{}
			default:
				http.Error(w, "Unknown entity for validation", http.StatusBadRequest)
				return
			}

			if err := json.NewDecoder(r.Body).Decode(instance); err != nil {
				http.Error(w, "Invalid JSON input", http.StatusBadRequest)
				return
			}

			if err := instance.Validate(); err != nil {
				http.Error(w, fmt.Sprintf("Validation error: %v", err), http.StatusUnprocessableEntity)
				return
			}

			// In the future for optimization
			// we can store validated instance in context for handler use

			next.ServeHTTP(w, r)
		})
	}
}
