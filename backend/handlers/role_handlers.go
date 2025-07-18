package handlers

import (
	"encoding/json"
	"net/http"
)

// ProfileHandler handles profile-related requests
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	roleHeader := r.Header.Get("X-User-Role")

	response := map[string]interface{}{
		"message":   "Profile access granted",
		"user_role": roleHeader,
		"endpoint":  "profile",
	}

	json.NewEncoder(w).Encode(response)
}

// HomePageHandler handles home page requests
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	roleHeader := r.Header.Get("X-User-Role")

	response := map[string]interface{}{
		"message":   "Welcome to home page",
		"user_role": roleHeader,
		"endpoint":  "home",
	}

	json.NewEncoder(w).Encode(response)
}

// MenuHandler handles menu requests
func MenuHandler(w http.ResponseWriter, r *http.Request) {
	roleHeader := r.Header.Get("X-User-Role")

	response := map[string]interface{}{
		"message":   "Menu access granted",
		"user_role": roleHeader,
		"endpoint":  "menu",
		"dishes":    []string{"Pizza", "Burger", "Pasta"}, // Mock data
	}

	json.NewEncoder(w).Encode(response)
}

// OrderDashboardHandler handles kitchen order dashboard
func OrderDashboardHandler(w http.ResponseWriter, r *http.Request) {
	roleHeader := r.Header.Get("X-User-Role")

	response := map[string]interface{}{
		"message":   "Order dashboard access granted",
		"user_role": roleHeader,
		"endpoint":  "order-dashboard",
		"orders":    []string{"Order #1", "Order #2"}, // Mock data
	}

	json.NewEncoder(w).Encode(response)
}

// DeliveryDashboardHandler handles delivery dashboard
func DeliveryDashboardHandler(w http.ResponseWriter, r *http.Request) {
	roleHeader := r.Header.Get("X-User-Role")

	response := map[string]interface{}{
		"message":    "Delivery dashboard access granted",
		"user_role":  roleHeader,
		"endpoint":   "delivery-dashboard",
		"deliveries": []string{"Delivery #1", "Delivery #2"}, // Mock data
	}

	json.NewEncoder(w).Encode(response)
}

// AdminHandler handles admin-only requests
func AdminHandler(w http.ResponseWriter, r *http.Request) {
	roleHeader := r.Header.Get("X-User-Role")

	response := map[string]interface{}{
		"message":    "Admin access granted",
		"user_role":  roleHeader,
		"endpoint":   "admin",
		"admin_data": "Sensitive admin information",
	}

	json.NewEncoder(w).Encode(response)
}
