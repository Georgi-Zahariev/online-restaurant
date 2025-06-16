package handlers

import (
	"log/slog"
	"net/http"
)

// /readyz
func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	slog.Debug("Readiness check hit")

	ready := true

	if ready {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("READY"))
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("NOT READY"))
	}
}
