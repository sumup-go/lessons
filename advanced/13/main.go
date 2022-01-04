//go:generate
package main

import (
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/segmentio/kafka-go"
	"log"
	"net/http"
)

func main() {
	e := echo.New()
	topic := "participants"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	srv := server{
		producer: conn,
	}

	RegisterHandlers(e, srv)

	e.Logger.Fatal(e.Start(":8080"))
}

type server struct {
	producer *kafka.Conn
}

func (s server) PostParticipant(ctx echo.Context) error {
	var req Participant
	if err := ctx.Bind(&req); err != nil {
		log.Fatal(err)
	}

	log.Printf("saving participant: %v", req)

	data, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	_, err = s.producer.WriteMessages(
		kafka.Message{
			Value: data,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	return ctx.NoContent(http.StatusCreated)
}
