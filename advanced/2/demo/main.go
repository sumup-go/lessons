package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type db struct{}

func (_ *db) GetData() string {
	return "Hello, World!"
}

func main() {
	d := &db{}

	r := mux.NewRouter()
	r.HandleFunc("/", handleIndex(d)).Methods(http.MethodGet)

	log.Println("starting server")
	log.Fatal(http.ListenAndServe(":8080", r))
}

type fetcher interface {
	GetData() string
}

func handleIndex(d fetcher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(d.GetData()))
	}
}
