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

type UserHandler struct {
	Manager *managers.Manager
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	logger := middlewares.GetLogger(r.Context())
	w.Header().Set("Content-Type", "application/json")
	users, err := h.Manager.GetAllUsers(r.Context())
	if err != nil {
		if logger != nil {
			logger.Error("GetAllUsers failed", "error", err)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	logger := middlewares.GetLogger(r.Context())
	w.Header().Set("Content-Type", "application/json")
	idStr := mux.Vars(r)["id"]
	if idStr == "" {
		if logger != nil {
			logger.Error("User ID is required")
		}
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	user, err := h.Manager.GetUser(r.Context(), idStr)
	if err != nil {
		if logger != nil {
			logger.Error("GetUser failed", "error", err)
		}
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
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
	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		if logger != nil {
			logger.Error("Invalid JSON input", "error", err)
		}
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}
	if err := h.Manager.CreateUser(r.Context(), &user); err != nil {
		if logger != nil {
			logger.Error("CreateUser failed", "error", err)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Fetch the created user to return DB-generated fields
	createdUser, _ := h.Manager.GetUser(r.Context(), user.ID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	logger := middlewares.GetLogger(r.Context())
	w.Header().Set("Content-Type", "application/json")
	idStr := mux.Vars(r)["id"]
	if idStr == "" {
		if logger != nil {
			logger.Error("User ID is required")
		}
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	var user models.User
	body, err := io.ReadAll(r.Body)
	if err != nil {
		if logger != nil {
			logger.Error("Failed to read request body", "error", err)
		}
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &user); err != nil {
		if logger != nil {
			logger.Error("Invalid JSON input", "error", err)
		}
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}
	if err := h.Manager.UpdateUser(r.Context(), idStr, &user); err != nil {
		if logger != nil {
			logger.Error("UpdateUser failed", "error", err)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	updatedUser, _ := h.Manager.GetUser(r.Context(), idStr)
	json.NewEncoder(w).Encode(updatedUser)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	logger := middlewares.GetLogger(r.Context())
	idStr := mux.Vars(r)["id"]
	if idStr == "" {
		if logger != nil {
			logger.Error("User ID is required")
		}
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	if err := h.Manager.DeleteUser(r.Context(), idStr); err != nil {
		if logger != nil {
			logger.Error("DeleteUser failed", "error", err)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// GetCurrentUser returns the current user from context
func (h *UserHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	logger := middlewares.GetLogger(r.Context())
	w.Header().Set("Content-Type", "application/json")

	user, err := h.Manager.GetCurrentUser(r.Context())
	if err != nil {
		if logger != nil {
			logger.Error("GetCurrentUser failed", "error", err)
		}
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}
