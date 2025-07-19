package middlewares

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gofrs/uuid/v5"
)

type ctxKeyLogger struct{}
type ctxKeyUser struct{}

var LoggerKey = ctxKeyLogger{}
var UserKey = ctxKeyUser{}

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

// Demoo for having user checking

// Middleware to extract user context from headers
// For now, we'll use a hardcoded user ID from X-User-ID header for testing
func UserContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// For demo purposes, we'll hardcode a user ID or get it from header
		userID := r.Header.Get("X-User-ID")
		if userID == "" {
			// Hardcoded user ID for testing - you can change this logic
			userID = "550e8400-e29b-41d4-a716-446655440001" // Demo user ID
		}

		ctx := context.WithValue(r.Context(), UserKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Demo to get user ID from context
// Helper to get user ID from context
func GetUserID(ctx context.Context) (string, bool) {
	if userID, ok := ctx.Value(UserKey).(string); ok {
		return userID, true
	}
	return "", false
}
