package main

import (
	"context"
	"log"

	"l11"

	"github.com/deepmap/oapi-codegen/pkg/types"
)

func main() {
	client, err := l11.NewClientWithResponses("http://localhost:3001")
	if err != nil {
		log.Fatal(err)
	}
	participant := l11.PostParticipantJSONRequestBody{
		Email: email("me@me.com"),
		Track: trackPointer(l11.ParticipantTrackBeginners),
	}

	response, err := client.PostParticipantWithResponse(context.Background(), participant)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(response.Status())

}

func trackPointer(track l11.ParticipantTrack) *l11.ParticipantTrack {
	return &track
}

func email(str string) *types.Email {
	e := types.Email(str)
	return &e
}
