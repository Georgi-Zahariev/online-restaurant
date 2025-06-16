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

	r.HandleFunc("/healthz", aroundHandlers(handlers.HealthHandler)).Methods("GET")
	r.HandleFunc("/readyz", aroundHandlers(handlers.ReadinessHandler)).Methods("GET")
	r.HandleFunc("/api/object1", aroundHandlers(handlers.Object1Handler)).Methods("GET")
	r.HandleFunc("/api/object2", aroundHandlers(handlers.Object2Handler)).Methods("GET")

	return r
}

func aroundHandlers(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//r.Context()
		log.Printf("DEBUG: Entering handler for %s %s", r.Method, r.URL.Path)
		next(w, r)

	}
}
