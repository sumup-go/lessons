//go:generate
package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	RegisterHandlers(e, server{})

	e.Logger.Fatal(e.Start(":8080"))
}

type server struct{}

func (s server) PostParticipant(ctx echo.Context) error {
	var req Participant
	if err := ctx.Bind(&req); err != nil {
		log.Fatal(err)
	}

	log.Printf("saving participant: %v", req)

	return ctx.NoContent(http.StatusCreated)
}
