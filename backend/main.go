package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Georgi-Zahariev/online-restaurant/backend/routers"
	"github.com/Georgi-Zahariev/online-restaurant/config"
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

	// Setup the router from the routers package
	r := routers.SetupRouter()

	// Start server
	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("INFO: Starting web server on port %d in %s mode...", cfg.Port, cfg.Env)
	log.Fatal(http.ListenAndServe(addr, r))
}
