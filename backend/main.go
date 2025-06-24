package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/Georgi-Zahariev/online-restaurant/backend/middlewares"
	"github.com/Georgi-Zahariev/online-restaurant/backend/server"
	"github.com/Georgi-Zahariev/online-restaurant/config"
	"github.com/sethvargo/go-envconfig"
)

func main() {
	// Load configuration
	var cfg config.Config
	ctx := context.Background()
	slog.Info("Application starting...")

	if err := envconfig.Process(ctx, &cfg); err != nil {
		slog.Error("Failed to load configuration", slog.String("error", err.Error()))
		os.Exit(1)
	}

	slog.Info("Configuration loaded successfully", slog.Int("Port", cfg.Port), slog.String("Env", cfg.Env))

	// Initialize the server
	s := &server.Server{}
	s.Initialize()

	// Apply middlewares
	s.Router.Use(middlewares.LoggerMiddleware)
	s.Router.Use(middlewares.JSONContentTypeMiddleware)
	//s.Router.Use(middlewares.AuthorizationMiddleware)

	// Start the server
	addr := fmt.Sprintf(":%d", cfg.Port)
	slog.Info("Starting web server", slog.String("address", addr), slog.String("environment", cfg.Env))

	if err := http.ListenAndServe(addr, s.Router); err != nil {
		slog.Error("Failed to start server", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
