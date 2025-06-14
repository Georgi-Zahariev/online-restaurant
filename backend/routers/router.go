package routers

import (
	"log"
	"net/http"

	"github.com/Georgi-Zahariev/online-restaurant/backend/handlers"
	"github.com/gorilla/mux"
)

// sets up the application routes.
func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/healthz", loggingMiddleware(handlers.HealthHandler)).Methods("GET")
	r.HandleFunc("/readyz", loggingMiddleware(handlers.ReadinessHandler)).Methods("GET")
	r.HandleFunc("/api/object1", loggingMiddleware(handlers.Object1Handler)).Methods("GET")
	r.HandleFunc("/api/object2", loggingMiddleware(handlers.Object2Handler)).Methods("GET")

	return r
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("DEBUG: Entering handler for %s %s", r.Method, r.URL.Path)
		next(w, r)
	}
}
