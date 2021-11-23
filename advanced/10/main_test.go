package main

import (
	"net/http"
	"testing"
	"time"
)

const timeout = time.Second * 2

func TestServerHandle(t *testing.T) {
	called := false

	sem := make(chan struct{})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1000 * time.Millisecond)
		called = true
		sem <- struct{}{}
	})

	go func() {
		http.ListenAndServe(":8080", nil)
	}()

	go func() {
		http.Get("http://localhost:8080/")
	}()

	select {
	case <-sem:
	case <-time.After(timeout):
		t.Fatal("timeout waiting for handler request")
	}

	if called != true {
		t.Fatal("called is false")
	}
}
