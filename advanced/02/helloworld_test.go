package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorldHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	handleHelloWorld(w, r)

	expectedBody := "Hello, World!"
	if w.Body.String() != expectedBody {
		t.Fatalf("unexpected body, expected: %s got: %s", expectedBody, w.Body.String())
	}

	if w.Code != http.StatusOK {
		t.Fatalf("unexpected status code, expected: %d got: %d", http.StatusOK, w.Code)
	}
}
