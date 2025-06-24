package middlewares

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gofrs/uuid/v5"
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

type ctxKeyLogger struct{}

var LoggerKey = ctxKeyLogger{}

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
