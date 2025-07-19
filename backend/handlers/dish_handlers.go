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

type DishHandler struct {
	Manager *managers.Manager
}

func (h *DishHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	logger := middlewares.GetLogger(r.Context())
	w.Header().Set("Content-Type", "application/json")
	dishes, err := h.Manager.GetAllDishes(r.Context())
	if err != nil {
		if logger != nil {
			logger.Error("GetAllDish failed", "error", err)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(dishes)
}

func (h *DishHandler) Get(w http.ResponseWriter, r *http.Request) {
	logger := middlewares.GetLogger(r.Context())
	w.Header().Set("Content-Type", "application/json")
	idStr := mux.Vars(r)["id"]
	if idStr == "" {
		if logger != nil {
			logger.Error("Dish ID is required")
		}
		http.Error(w, "Dish ID is required", http.StatusBadRequest)
		return
	}
	dish, err := h.Manager.GetDish(r.Context(), idStr)
	if err != nil {
		if logger != nil {
			logger.Error("GetDish failed", "error", err)
		}
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(dish)
}

func (h *DishHandler) Create(w http.ResponseWriter, r *http.Request) {
	logger := middlewares.GetLogger(r.Context())
	w.Header().Set("Content-Type", "application/json")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		if logger != nil {
			logger.Error("Failed to read request body", "error", err)
		}
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	var dish models.Dish
	if err := json.Unmarshal(body, &dish); err != nil {
		if logger != nil {
			logger.Error("Invalid JSON input", "error", err)
		}
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}
	if err := h.Manager.CreateDish(r.Context(), &dish); err != nil {
		if logger != nil {
			logger.Error("CreateDish failed", "error", err)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	createdDish, _ := h.Manager.GetDish(r.Context(), dish.ID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdDish)
}

func (h *DishHandler) Update(w http.ResponseWriter, r *http.Request) {
	logger := middlewares.GetLogger(r.Context())
	w.Header().Set("Content-Type", "application/json")
	idStr := mux.Vars(r)["id"]
	if idStr == "" {
		if logger != nil {
			logger.Error("Dish ID is required")
		}
		http.Error(w, "Dish ID is required", http.StatusBadRequest)
		return
	}
	var dish models.Dish
	body, err := io.ReadAll(r.Body)
	if err != nil {
		if logger != nil {
			logger.Error("Failed to read request body", "error", err)
		}
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &dish); err != nil {
		if logger != nil {
			logger.Error("Invalid JSON input", "error", err)
		}
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}
	if err := h.Manager.UpdateDish(r.Context(), idStr, &dish); err != nil {
		if logger != nil {
			logger.Error("UpdateDish failed", "error", err)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	updatedDish, _ := h.Manager.GetDish(r.Context(), idStr)
	json.NewEncoder(w).Encode(updatedDish)
}

func (h *DishHandler) Delete(w http.ResponseWriter, r *http.Request) {
	logger := middlewares.GetLogger(r.Context())
	idStr := mux.Vars(r)["id"]
	if idStr == "" {
		if logger != nil {
			logger.Error("Dish ID is required")
		}
		http.Error(w, "Dish ID is required", http.StatusBadRequest)
		return
	}
	if err := h.Manager.DeleteDish(r.Context(), idStr); err != nil {
		if logger != nil {
			logger.Error("DeleteDish failed", "error", err)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
