package handlers

import (
	"log/slog"
	"net/http"
)

// /healthz

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	slog.Debug("Health check hit")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
