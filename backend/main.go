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

	if err := envconfig.Process(ctx, &cfg); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/api/object1", object1Handler).Methods("GET")
	r.HandleFunc("/api/object2", object2Handler).Methods("GET")
	r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	r.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ready"))
	})

	addr := fmt.Sprintf(":%d", cfg.Port)
	fmt.Printf("App starting on port %d in %s mode\n", cfg.Port, cfg.Env)
	log.Fatal(http.ListenAndServe(addr, r))
}
