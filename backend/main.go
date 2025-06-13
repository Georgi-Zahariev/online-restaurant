package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Georgi-Zahariev/online-restaurant/config"
	"github.com/gorilla/mux"
	"github.com/sethvargo/go-envconfig"
)

func main() {
	var cfg config.Config
	ctx := context.Background()

	log.Println("INFO: Starting application...")

	// Load configuration
	if err := envconfig.Process(ctx, &cfg); err != nil {
		log.Fatalf("ERROR: Failed to load configuration: %v", err)
	}
	log.Println("INFO: Configuration loaded successfully.")
	log.Printf("DEBUG: Configuration values: Port=%d, Env=%s, LogFormat=%s, LogSeverity=%s", cfg.Port, cfg.Env, cfg.LogFormat, cfg.LogLevel)

	// Create router
	r := mux.NewRouter()
	r.HandleFunc("/api/object1", loggingMiddleware(handlers.object1Handler)).Methods("GET")
	r.HandleFunc("/api/object2", loggingMiddleware(handlers.object2Handler)).Methods("GET")
	r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		log.Println("DEBUG: Entering healthz handler.")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	r.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		log.Println("DEBUG: Entering readyz handler.")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ready"))
	})

	// Start server
	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("INFO: Starting web server on port %d in %s mode...", cfg.Port, cfg.Env)
	log.Fatal(http.ListenAndServe(addr, r))
}

// Middleware to log debug messages when entering handlers
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("DEBUG: Entering handler for %s %s", r.Method, r.URL.Path)
		next(w, r)
	}
}
