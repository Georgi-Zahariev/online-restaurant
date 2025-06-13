package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestObject1Handler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/object1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Object1Handler)
	handler.ServeHTTP(rr, req)

	// Check status code.
	if rr.Code != http.StatusOK {
		t.Errorf("object1Handler: expected status %d, got %d", http.StatusOK, rr.Code)
	}

	// Check body.
	expected := "object1"
	if rr.Body.String() != expected {
		t.Errorf("object1Handler: expected body %q, got %q", expected, rr.Body.String())
	}
}

func TestObject2Handler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/object2", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Object2Handler)
	handler.ServeHTTP(rr, req)

	// Check status code.
	if rr.Code != http.StatusOK {
		t.Errorf("object2Handler: expected status %d, got %d", http.StatusOK, rr.Code)
	}

	// Check body.
	expected := "object2"
	if rr.Body.String() != expected {
		t.Errorf("object2Handler: expected body %q, got %q", expected, rr.Body.String())
	}
}

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/healthz", nil)
	rr := httptest.NewRecorder()

	HealthHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("HealthHandler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "OK"
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("HealthHandler returned unexpected body: got %q want %q", rr.Body.String(), expected)
	}
}

func TestReadinessHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/readyz", nil)
	rr := httptest.NewRecorder()

	ReadinessHandler(rr, req)

	// expected status is 200 (ready).
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("ReadinessHandler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "READY"
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("ReadinessHandler returned unexpected body: got %q want %q", rr.Body.String(), expected)
	}
}
