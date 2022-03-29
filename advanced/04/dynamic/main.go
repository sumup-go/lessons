package main

import (
	"log"
	"net/http"

	_ "embed"

	"github.com/gorilla/mux"
)

var (
	//go:embed index.html
	indexFile []byte
)

func main() {
	d := newDB()

	r := mux.NewRouter()
	r.HandleFunc("/", index(d)).Methods(http.MethodGet)

	log.Println("starting server on port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

//go:generate mockgen -source=$GOFILE -package=main -destination=main_mock.go
type dber interface {
	GetData() []byte
}

func index(d dber) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write(d.GetData())
	}
}
