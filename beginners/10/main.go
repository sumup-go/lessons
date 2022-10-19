package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/my-page", indexHandler)

	fmt.Println("starting server on port: 8080")

	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	b, err := os.ReadFile("index1.html")
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		log.Println(err)
	}
}
