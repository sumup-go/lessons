package main

import (
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
)

func static(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	m := mux.NewRouter()
	m.HandleFunc("/static", static)
	
	log.Fatal(http.ListenAndServe(":9999", m))
}
