package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sumup-go/lessons/advanced/12/sqlc"
)

type DB struct {
	*sqlc.Queries
}

func New() (*DB, error) {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("connecting to postgres: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("pinging db: %v", err)
	}

	queries := sqlc.New(db)

	return &DB{
		Queries: queries,
	}, nil
}
