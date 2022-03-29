package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleIndex(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	handleIndex(&db{})(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedBody := "Hello, World!"
	assert.Equal(t, expectedBody, w.Body.String())
}
