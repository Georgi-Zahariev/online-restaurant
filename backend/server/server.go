package server

import (
	"log"
	"net/http"

	"github.com/Georgi-Zahariev/online-restaurant/backend/handlers"
	"github.com/Georgi-Zahariev/online-restaurant/backend/managers"
	"github.com/Georgi-Zahariev/online-restaurant/backend/routers"
	"github.com/gorilla/mux"
)

type Server struct {
	Router  *mux.Router
	Manager *managers.Manager
}

func (s *Server) Initialize() {
	s.Manager = managers.NewManager()
	s.Router = mux.NewRouter()

	// Initialize generic handlers
	userHandler := &handlers.GenericHandler{Manager: s.Manager, Entity: "users"}
	dishHandler := &handlers.GenericHandler{Manager: s.Manager, Entity: "dishes"}
	orderHandler := &handlers.GenericHandler{Manager: s.Manager, Entity: "orders"}

	// Register routes
	routers.InitializeRoutes(s.Router, map[string]map[string]http.HandlerFunc{
		"users": {
			"GET":    userHandler.GetAll,
			"POST":   userHandler.Create,
			"GET_ID": userHandler.Get,
			"PUT":    userHandler.Update,
			"DELETE": userHandler.Delete,
		},
		"dishes": {
			"GET":    dishHandler.GetAll,
			"POST":   dishHandler.Create,
			"GET_ID": dishHandler.Get,
			"PUT":    dishHandler.Update,
			"DELETE": dishHandler.Delete,
		},
		"orders": {
			"GET":    orderHandler.GetAll,
			"POST":   orderHandler.Create,
			"GET_ID": orderHandler.Get,
			"PUT":    orderHandler.Update,
			"DELETE": orderHandler.Delete,
		},
	})
}

func (s *Server) Run(addr string) {
	log.Printf("Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, s.Router))
}
