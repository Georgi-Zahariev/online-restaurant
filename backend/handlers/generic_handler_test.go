package handlers

/*import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Georgi-Zahariev/online-restaurant/backend/managers"
	"github.com/gorilla/mux"
)

func setupTestHandler(entity string) *GenericHandler {
	manager := managers.NewManager()
	return &GenericHandler{Manager: manager, Entity: entity}
}

func setupRouter(handler *GenericHandler, method, path string, handlerFunc http.HandlerFunc) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(path, handlerFunc).Methods(method)
	return router
}

func executeRequest(router *mux.Router, method, path string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr
}

func TestGetAll(t *testing.T) {
	handler := setupTestHandler("users")
	router := setupRouter(handler, "GET", "/api/users", handler.GetAll)

	rr := executeRequest(router, "GET", "/api/users", nil)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var actual []map[string]interface{}
	expected := []map[string]interface{}{
		{"id": "1", "phone_number": "123456789"},
		{"id": "2", "phone_number": "987654321"},
	}

	if err := json.Unmarshal(rr.Body.Bytes(), &actual); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if len(actual) != len(expected) {
		t.Errorf("Expected %d items, got %d", len(expected), len(actual))
	}

	for i, item := range actual {
		if item["id"] != expected[i]["id"] || item["phone_number"] != expected[i]["phone_number"] {
			t.Errorf("Expected item %v, got %v", expected[i], item)
		}
	}
}

func TestGet(t *testing.T) {
	handler := setupTestHandler("users")
	router := setupRouter(handler, "GET", "/api/users/{id}", handler.Get)

	rr := executeRequest(router, "GET", "/api/users/1", nil)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var actual map[string]interface{}
	expected := map[string]interface{}{
		"id":           "1",
		"phone_number": "123456789",
	}

	if err := json.Unmarshal(rr.Body.Bytes(), &actual); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if actual["id"] != expected["id"] || actual["phone_number"] != expected["phone_number"] {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestGet_NotFound(t *testing.T) {
	handler := setupTestHandler("users")
	router := setupRouter(handler, "GET", "/api/users/{id}", handler.Get)

	rr := executeRequest(router, "GET", "/api/users/999", nil)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, rr.Code)
	}
}

func TestCreate(t *testing.T) {
	handler := setupTestHandler("users")
	router := setupRouter(handler, "POST", "/api/users", handler.Create)

	body := []byte(`{"id":"3","phone_number":"555555555"}`)
	rr := executeRequest(router, "POST", "/api/users", body)

	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, rr.Code)
	}

	var actual map[string]interface{}
	expected := map[string]interface{}{
		"id":           "3",
		"phone_number": "555555555",
	}

	if err := json.Unmarshal(rr.Body.Bytes(), &actual); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if actual["id"] != expected["id"] || actual["phone_number"] != expected["phone_number"] {
		t.Errorf("Expected body %v, got %v", expected, actual)
	}
}

func TestCreate_InvalidJSON(t *testing.T) {
	handler := setupTestHandler("users")
	router := setupRouter(handler, "POST", "/api/users", handler.Create)

	body := []byte(`invalid-json`)
	rr := executeRequest(router, "POST", "/api/users", body)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rr.Code)
	}
}

func TestUpdate(t *testing.T) {
	handler := setupTestHandler("users")
	router := setupRouter(handler, "PUT", "/api/users/{id}", handler.Update)

	body := []byte(`{"id":"1","phone_number":"111111111"}`)
	rr := executeRequest(router, "PUT", "/api/users/1", body)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var actual map[string]interface{}
	expected := map[string]interface{}{
		"id":           "1",
		"phone_number": "111111111",
	}

	if err := json.Unmarshal(rr.Body.Bytes(), &actual); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if actual["id"] != expected["id"] || actual["phone_number"] != expected["phone_number"] {
		t.Errorf("Expected body %v, got %v", expected, actual)
	}
}

func TestUpdate_InvalidJSON(t *testing.T) {
	handler := setupTestHandler("users")
	router := setupRouter(handler, "PUT", "/api/users/{id}", handler.Update)

	body := []byte(`invalid-json`)
	rr := executeRequest(router, "PUT", "/api/users/1", body)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rr.Code)
	}
}

func TestDelete(t *testing.T) {
	handler := setupTestHandler("users")
	router := setupRouter(handler, "DELETE", "/api/users/{id}", handler.Delete)

	rr := executeRequest(router, "DELETE", "/api/users/1", nil)

	if rr.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, got %d", http.StatusNoContent, rr.Code)
	}
}

func TestDelete_NotFound(t *testing.T) {
	handler := setupTestHandler("users")
	router := setupRouter(handler, "DELETE", "/api/users/{id}", handler.Delete)

	rr := executeRequest(router, "DELETE", "/api/users/999", nil)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, rr.Code)
	}
}
*/
