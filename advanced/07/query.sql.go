// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package main

import (
	"context"
	"database/sql"
)

const createParticipant = `-- name: CreateParticipant :exec
INSERT INTO participant(email, track, name) VALUES ($1, $2, $3)
`

type CreateParticipantParams struct {
	Email string
	Track string
	Name  sql.NullString
}

func (q *Queries) CreateParticipant(ctx context.Context, arg CreateParticipantParams) error {
	_, err := q.db.ExecContext(ctx, createParticipant, arg.Email, arg.Track, arg.Name)
	return err
}

const getParticipant = `-- name: GetParticipant :one
SELECT id, email, track, name FROM participant WHERE email = $1
`

func (q *Queries) GetParticipant(ctx context.Context, email string) (Participant, error) {
	row := q.db.QueryRowContext(ctx, getParticipant, email)
	var i Participant
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Track,
		&i.Name,
	)
	return i, err
}
