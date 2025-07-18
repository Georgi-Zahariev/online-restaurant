package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Georgi-Zahariev/online-restaurant/backend/models"
	"github.com/gofrs/uuid/v5"
)

type ctxKeyLogger struct{}

var LoggerKey = ctxKeyLogger{}

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

// Simple role-based access middleware
// Expects role to be passed via X-User-Role header (temporary solution until IDP is implemented)
func RequireRole(allowedRoles ...models.UserRole) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get role from header (temporary solution)
			roleHeader := r.Header.Get("X-User-Role")
			if roleHeader == "" {
				w.WriteHeader(http.StatusBadRequest)
				response := map[string]string{"error": "Missing X-User-Role header"}
				json.NewEncoder(w).Encode(response)
				return
			}

			userRole := models.UserRole(roleHeader)

			// Owner can access everything
			if userRole == models.RoleOwner {
				next.ServeHTTP(w, r)
				return
			}

			// Check if user role is in allowed roles
			for _, role := range allowedRoles {
				if userRole == role {
					next.ServeHTTP(w, r)
					return
				}
			}

			w.WriteHeader(http.StatusForbidden)
			response := map[string]string{
				"error":          "Access denied - insufficient role",
				"user_role":      string(userRole),
				"required_roles": fmt.Sprintf("%v", allowedRoles),
			}
			json.NewEncoder(w).Encode(response)
		})
	}
}

// Middleware to add a logger with correlation ID to the context
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := uuid.Must(uuid.NewV4()).String()
		logger := slog.Default().With("correlation_id", reqID)
		ctx := context.WithValue(r.Context(), LoggerKey, logger)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Helper to get logger from context
func GetLogger(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(LoggerKey).(*slog.Logger); ok {
		return logger
	}
	return slog.Default()
}
