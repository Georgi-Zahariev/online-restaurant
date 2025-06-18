package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Georgi-Zahariev/online-restaurant/backend/managers"
	"github.com/gorilla/mux"
)

type GenericHandler struct {
	Manager *managers.Manager
	Entity  string // The entity this handler is responsible for (e.g., "users", "dishes", "orders")
}

func (h *GenericHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.Manager.GetAll(h.Entity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (h *GenericHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	data, err := h.Manager.Get(h.Entity, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (h *GenericHandler) Create(w http.ResponseWriter, r *http.Request) {
	var instance interface{}
	if err := json.NewDecoder(r.Body).Decode(&instance); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}
	if err := h.Manager.Create(h.Entity, instance); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(instance)
}

func (h *GenericHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var instance interface{}
	if err := json.NewDecoder(r.Body).Decode(&instance); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}
	if err := h.Manager.Update(h.Entity, id, instance); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(instance)
}

func (h *GenericHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := h.Manager.Delete(h.Entity, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
