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

type OrderHandler struct {
	Manager *managers.Manager
}

func (h *OrderHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	logger := middlewares.GetLogger(r.Context())
	w.Header().Set("Content-Type", "application/json")
	orders, err := h.Manager.GetAllOrder()
	if err != nil {
		if logger != nil {
			logger.Error("GetAllOrder failed", "error", err)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(orders)
}

func (h *OrderHandler) Get(w http.ResponseWriter, r *http.Request) {
	logger := middlewares.GetLogger(r.Context())
	w.Header().Set("Content-Type", "application/json")
	idStr := mux.Vars(r)["id"]
	if idStr == "" {
		if logger != nil {
			logger.Error("Order ID is required")
		}
		http.Error(w, "Order ID is required", http.StatusBadRequest)
		return
	}
	order, err := h.Manager.GetOrder(idStr)
	if err != nil {
		if logger != nil {
			logger.Error("GetOrder failed", "error", err)
		}
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
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
	var order models.Order
	if err := json.Unmarshal(body, &order); err != nil {
		if logger != nil {
			logger.Error("Invalid JSON input", "error", err)
		}
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}
	if err := h.Manager.CreateOrder(&order); err != nil {
		if logger != nil {
			logger.Error("CreateOrder failed", "error", err)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	createdOrder, _ := h.Manager.GetOrder(order.ID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdOrder)
}

func (h *OrderHandler) Update(w http.ResponseWriter, r *http.Request) {
	logger := middlewares.GetLogger(r.Context())
	w.Header().Set("Content-Type", "application/json")
	idStr := mux.Vars(r)["id"]
	if idStr == "" {
		if logger != nil {
			logger.Error("Order ID is required")
		}
		http.Error(w, "Order ID is required", http.StatusBadRequest)
		return
	}
	var order models.Order
	body, err := io.ReadAll(r.Body)
	if err != nil {
		if logger != nil {
			logger.Error("Failed to read request body", "error", err)
		}
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &order); err != nil {
		if logger != nil {
			logger.Error("Invalid JSON input", "error", err)
		}
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}
	if err := h.Manager.UpdateOrder(idStr, &order); err != nil {
		if logger != nil {
			logger.Error("UpdateOrder failed", "error", err)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	updatedOrder, _ := h.Manager.GetOrder(idStr)
	json.NewEncoder(w).Encode(updatedOrder)
}

func (h *OrderHandler) Delete(w http.ResponseWriter, r *http.Request) {
	logger := middlewares.GetLogger(r.Context())
	idStr := mux.Vars(r)["id"]
	if idStr == "" {
		if logger != nil {
			logger.Error("Order ID is required")
		}
		http.Error(w, "Order ID is required", http.StatusBadRequest)
		return
	}
	if err := h.Manager.DeleteOrder(idStr); err != nil {
		if logger != nil {
			logger.Error("DeleteOrder failed", "error", err)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
