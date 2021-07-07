package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleHelloWorld)

	log.Println("starting server")
	http.ListenAndServe(":8080", nil)
}

func handleHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}
