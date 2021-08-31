package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	//nolint:deadcode,unused,varcheck
	something = ""
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Hello, World\n")); err != nil {
		log.Println("err: failed to write response")
		http.Error(w, "failed to write response", http.StatusInternalServerError)

		return
	}
}

func main() {
	e := echo.New()
	e.GET("/", echo.WrapHandler(http.HandlerFunc(indexHandler)))

	e.Logger.Fatal(e.Start(":8080"))
}
