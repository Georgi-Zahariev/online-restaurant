package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/Georgi-Zahariev/online-restaurant/backend/middlewares"
	"github.com/Georgi-Zahariev/online-restaurant/backend/routers"
	"github.com/Georgi-Zahariev/online-restaurant/config"
	"github.com/sethvargo/go-envconfig"
)

func main() {
	var cfg config.Config
	ctx := context.Background()
	slog.Info("Application starting...")

	// Load configuration
	if err := envconfig.Process(ctx, &cfg); err != nil {
		slog.Error("Failed to load configuration", slog.String("error", err.Error()))
	}

	slog.Debug("Loaded AUTH_TOKEN", slog.String("AuthToken", cfg.AuthToken)) // Log the token

	slog.Info("Configuration loaded successfully")
	slog.Debug("Configuration values", slog.Int("Port", cfg.Port), slog.String("Env", cfg.Env), slog.String("LogFormat", cfg.LogFormat), slog.String("LogLevel", cfg.LogLevel))

	// Setup the router from the routers package
	r := routers.SetupRouter()

	// Apply middlewares
	r.Use(middlewares.JSONContentTypeMiddleware)
	r.Use(middlewares.AuthorizationMiddleware(cfg.AuthToken)) // Pass the expected token

	// Example endpoint
	r.HandleFunc("/example", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{"message": "Success"}
		json.NewEncoder(w).Encode(response)
	}).Methods("GET")

	// Start server
	addr := fmt.Sprintf(":%d", cfg.Port)
	slog.Info("Starting web server", slog.String("address", addr), slog.String("environment", cfg.Env))

	// log.Fatal(http.ListenAndServe(addr, r))
	err := http.ListenAndServe(addr, r)
	if err != nil {
		slog.Error("Failed to start server", slog.String("address", addr), slog.String("error", err.Error()))
		// same behavior as log.Fatal, but without exiting the program
		os.Exit(1)
	}
}
