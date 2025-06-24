package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Georgi-Zahariev/online-restaurant/backend/managers"
	"github.com/Georgi-Zahariev/online-restaurant/backend/middlewares"
	"github.com/Georgi-Zahariev/online-restaurant/backend/models"
	"github.com/gorilla/mux"
)

type GenericHandler struct {
	Manager *managers.Manager
	Entity  string // The entity this handler is responsible for (e.g., "users", "dishes", "orders")
}

func (h *GenericHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	logger := middlewares.GetLogger(r.Context())
	logger.Info("GetAll called", "entity", h.Entity)
	data, err := h.Manager.GetAll(h.Entity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (h *GenericHandler) Get(w http.ResponseWriter, r *http.Request) {
	logger := middlewares.GetLogger(r.Context())
	id := mux.Vars(r)["id"]
	logger.Info("Get called", "entity", h.Entity, "id", id)
	data, err := h.Manager.Get(h.Entity, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (h *GenericHandler) Create(w http.ResponseWriter, r *http.Request) {
	logger := middlewares.GetLogger(r.Context())
	logger.Info("Create called", "entity", h.Entity)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	var instance interface{}

	switch h.Entity {
	case "users":
		var user models.User
		if err := json.Unmarshal(body, &user); err != nil {
			logger.Error("Invalid JSON input", "error", err)
			http.Error(w, "Invalid JSON input", http.StatusBadRequest)
			return
		}
		instance = user
	case "dishes":
		var dish models.Dish
		if err := json.Unmarshal(body, &dish); err != nil {
			logger.Error("Invalid JSON input", "error", err)
			http.Error(w, "Invalid JSON input", http.StatusBadRequest)
			return
		}
		instance = dish
	case "orders":
		var order models.Order
		if err := json.Unmarshal(body, &order); err != nil {
			logger.Error("Invalid JSON input", "error", err)
			http.Error(w, "Invalid JSON input", http.StatusBadRequest)
			return
		}
		instance = order
	default:
		logger.Error("Unknown entity", "entity", h.Entity)
		http.Error(w, "Unknown entity", http.StatusBadRequest)
		return
	}

	if err := h.Manager.Create(h.Entity, instance); err != nil {
		logger.Error("Create failed", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response, _ := json.Marshal(instance)
	w.Write(response)

	logger.Info("Create successful", "entity", h.Entity)
}

func (h *GenericHandler) Update(w http.ResponseWriter, r *http.Request) {
	logger := middlewares.GetLogger(r.Context())
	id := mux.Vars(r)["id"]
	logger.Info("Update called", "entity", h.Entity, "id", id)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error("Failed to read request body", "error", err)
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	var instance interface{}

	switch h.Entity {
	case "users":
		var user models.User
		if err := json.Unmarshal(body, &user); err != nil {
			logger.Error("Invalid JSON input", "error", err)
			http.Error(w, "Invalid JSON input", http.StatusBadRequest)
			return
		}
		instance = user
	case "dishes":
		var dish models.Dish
		if err := json.Unmarshal(body, &dish); err != nil {
			logger.Error("Invalid JSON input", "error", err)
			http.Error(w, "Invalid JSON input", http.StatusBadRequest)
			return
		}
		instance = dish
	case "orders":
		var order models.Order
		if err := json.Unmarshal(body, &order); err != nil {
			logger.Error("Invalid JSON input", "error", err)
			http.Error(w, "Invalid JSON input", http.StatusBadRequest)
			return
		}
		instance = order
	default:
		logger.Error("Unknown entity", "entity", h.Entity)
		http.Error(w, "Unknown entity", http.StatusBadRequest)
		return
	}

	if err := h.Manager.Update(h.Entity, id, instance); err != nil {
		logger.Error("Update failed", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(instance)
	w.Write(response)
	logger.Info("Update successful", "entity", h.Entity, "id", id)
}

func (h *GenericHandler) Delete(w http.ResponseWriter, r *http.Request) {
	logger := middlewares.GetLogger(r.Context())
	id := mux.Vars(r)["id"]
	logger.Info("Delete called", "entity", h.Entity, "id", id)
	if err := h.Manager.Delete(h.Entity, id); err != nil {
		logger.Error("Delete failed", "error", err, "id", id)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	logger.Info("Delete successful", "entity", h.Entity, "id", id)
}
