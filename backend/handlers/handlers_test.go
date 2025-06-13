package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestObject1Handler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/object1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(object1Handler)
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

	handler := http.HandlerFunc(object2Handler)
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
