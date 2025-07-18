package routers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Georgi-Zahariev/online-restaurant/backend/handlers"
	"github.com/Georgi-Zahariev/online-restaurant/backend/middlewares"
	"github.com/Georgi-Zahariev/online-restaurant/backend/models"
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

	// Public routes (no role checking required)
	r.HandleFunc("/healthz", aroundHandlers(handlers.HealthHandler)).Methods("GET")
	r.HandleFunc("/readyz", aroundHandlers(handlers.ReadinessHandler)).Methods("GET")

	// API routes with role-based access control
	api := r.PathPrefix("/api").Subrouter()

	// Customer routes (Home, Menu, Shopping cart, Profile)
	api.Handle("/home", middlewares.RequireRole(models.RoleCustomer)(http.HandlerFunc(handlers.HomePageHandler))).Methods("GET")
	api.Handle("/menu", middlewares.RequireRole(models.RoleCustomer)(http.HandlerFunc(handlers.MenuHandler))).Methods("GET")
	api.Handle("/dishes", middlewares.RequireRole(models.RoleCustomer)(http.HandlerFunc(handlers.Object1Handler))).Methods("GET")
	api.Handle("/cart", middlewares.RequireRole(models.RoleCustomer)(http.HandlerFunc(handlers.Object1Handler))).Methods("GET", "POST")
	api.Handle("/orders", middlewares.RequireRole(models.RoleCustomer)(http.HandlerFunc(handlers.Object2Handler))).Methods("GET", "POST")

	// Kitchen routes (Profile, Order dashboard)
	api.Handle("/order-dashboard", middlewares.RequireRole(models.RoleKitchen)(http.HandlerFunc(handlers.OrderDashboardHandler))).Methods("GET")
	api.Handle("/order-items/complete", middlewares.RequireRole(models.RoleKitchen)(http.HandlerFunc(handlers.Object2Handler))).Methods("PUT")

	// Delivery routes (Profile, Delivery dashboard)
	api.Handle("/delivery-dashboard", middlewares.RequireRole(models.RoleDelivery)(http.HandlerFunc(handlers.DeliveryDashboardHandler))).Methods("GET")
	api.Handle("/orders/deliver", middlewares.RequireRole(models.RoleDelivery)(http.HandlerFunc(handlers.Object1Handler))).Methods("PUT")

	// Profile routes (accessible by all roles)
	api.Handle("/profile", middlewares.RequireRole(models.RoleCustomer, models.RoleKitchen, models.RoleDelivery, models.RoleOwner)(http.HandlerFunc(handlers.ProfileHandler))).Methods("GET", "PUT")

	// Owner routes (can access everything - these are just examples)
	api.Handle("/admin", middlewares.RequireRole(models.RoleOwner)(http.HandlerFunc(handlers.AdminHandler))).Methods("GET")
	api.Handle("/analytics", middlewares.RequireRole(models.RoleOwner)(http.HandlerFunc(handlers.Object2Handler))).Methods("GET")

	return r
}

func aroundHandlers(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//r.Context()
		slog.Debug("Entering handler", slog.String("method", r.Method), slog.String("path", r.URL.Path))
		next(w, r)
	}
}
