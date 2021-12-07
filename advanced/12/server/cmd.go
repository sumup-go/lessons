package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"l11"
)

type server struct{}

func (s server) PostParticipant(ctx echo.Context) error {
	var req l11.Participant
	err := ctx.Bind(&req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("saving participant: %v", req)

	return ctx.NoContent(http.StatusCreated)
}

func main() {
	e := echo.New()
	l11.RegisterHandlers(e, server{})

	e.Logger.Fatal(e.Start(":3001"))
}
