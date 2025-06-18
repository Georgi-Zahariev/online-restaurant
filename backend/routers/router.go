package routers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Georgi-Zahariev/online-restaurant/backend/handlers"
	"github.com/gorilla/mux"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

func InitializeRoutes(router *mux.Router, handlers map[string]map[string]http.HandlerFunc) {
	for resource, methods := range handlers {
		for method, handler := range methods {
			switch method {
			case "GET":
				router.HandleFunc(fmt.Sprintf("/api/%s", resource), handler).Methods(GET)
			case "POST":
				router.HandleFunc(fmt.Sprintf("/api/%s", resource), handler).Methods(POST)
			case "GET_ID":
				router.HandleFunc(fmt.Sprintf("/api/%s/{id}", resource), handler).Methods(GET)
			case "PUT":
				router.HandleFunc(fmt.Sprintf("/api/%s/{id}", resource), handler).Methods(PUT)
			case "DELETE":
				router.HandleFunc(fmt.Sprintf("/api/%s/{id}", resource), handler).Methods(DELETE)
			}
		}
	}
}

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
		slog.Debug("Entering handler", slog.String("method", r.Method), slog.String("path", r.URL.Path))
		next(w, r)
	}
}
