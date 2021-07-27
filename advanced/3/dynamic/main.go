package main

import (
	"html/template"
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
)

type Page struct {
	Name string
}

func submission(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("index.html"))
	
	if r.Method != http.MethodPost {
		tmp.Execute(w, nil)
		return
	}
	
	name := r.FormValue("name")
	log.Println("received name", name)
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	p := Page{Name: name}
	tmp.Execute(w, p)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/submission", submission).Methods(http.MethodGet, http.MethodPost)
	
	log.Fatal(http.ListenAndServe(":9098", router))
}
