package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Georgi-Zahariev/online-restaurant/backend/handlers"
	"github.com/Georgi-Zahariev/online-restaurant/backend/managers"
	"github.com/Georgi-Zahariev/online-restaurant/backend/middlewares"
	"github.com/Georgi-Zahariev/online-restaurant/config"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	Router  *mux.Router
	Manager *managers.Manager
	DB      *gorm.DB
}

func (s *Server) Initialize(cfg config.Config) {
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	s.DB = db

	// Note: We're using manual migrations instead of AutoMigrate
	// Run migrations using docker-compose migrations service or manually
	// if err := db.AutoMigrate(&models.User{}, &models.Dish{}, &models.Order{}); err != nil {
	//     log.Fatalf("migration failed: %v", err)
	// }

	s.Manager = managers.NewManager(db)

	//
	s.Router = mux.NewRouter()

	// Apply user context middleware to all routes
	s.Router.Use(middlewares.UserContextMiddleware)

	// Initialize entity-specific handlers
	userHandler := &handlers.UserHandler{Manager: s.Manager}
	dishHandler := &handlers.DishHandler{Manager: s.Manager}
	orderHandler := &handlers.OrderHandler{Manager: s.Manager}

	// Register user routes
	s.Router.HandleFunc("/api/users", userHandler.GetAll).Methods("GET")
	s.Router.HandleFunc("/api/users/me", userHandler.GetCurrentUser).Methods("GET")
	s.Router.HandleFunc("/api/users/{id}", userHandler.Get).Methods("GET")
	s.Router.HandleFunc("/api/users", userHandler.Create).Methods("POST")
	s.Router.HandleFunc("/api/users/{id}", userHandler.Update).Methods("PUT")
	s.Router.HandleFunc("/api/users/{id}", userHandler.Delete).Methods("DELETE")

	// Register dish routes
	s.Router.HandleFunc("/api/dishes", dishHandler.GetAll).Methods("GET")
	s.Router.HandleFunc("/api/dishes/{id}", dishHandler.Get).Methods("GET")
	s.Router.HandleFunc("/api/dishes", dishHandler.Create).Methods("POST")
	s.Router.HandleFunc("/api/dishes/{id}", dishHandler.Update).Methods("PUT")
	s.Router.HandleFunc("/api/dishes/{id}", dishHandler.Delete).Methods("DELETE")

	// Register order routes
	s.Router.HandleFunc("/api/orders", orderHandler.GetAll).Methods("GET")
	s.Router.HandleFunc("/api/orders/{id}", orderHandler.Get).Methods("GET")
	s.Router.HandleFunc("/api/orders", orderHandler.Create).Methods("POST")
	s.Router.HandleFunc("/api/orders/{id}", orderHandler.Update).Methods("PUT")
	s.Router.HandleFunc("/api/orders/{id}", orderHandler.Delete).Methods("DELETE")

}
func (s *Server) Run(addr string) {
	log.Printf("Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, s.Router))
}
