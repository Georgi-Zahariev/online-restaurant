package handlers

import (
	"log/slog"
	"net/http"
)

func object1Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("object1"))
}

func object2Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("object2"))
}

// /healthz
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	slog.Debug("Health check hit")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

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
