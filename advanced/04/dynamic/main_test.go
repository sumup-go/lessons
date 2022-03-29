package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	var (
		payload = []byte("To the moon, Morty.")
	)

	ctrl := gomock.NewController(t)
	d := NewMockdber(ctrl)
	d.EXPECT().GetData().Return(payload)

	index(d)(w, r)

	assert.Equal(t, w.Code, 200)
	assert.Equal(t, payload, w.Body.Bytes())
}
