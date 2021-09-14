package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func migrateDB() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	mPath := "file://migrations"

	migrator, err := migrate.NewWithDatabaseInstance(mPath, "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	if err = migrator.Up(); err != nil && err.Error() != "no change" {
		log.Fatal(err)
	}
}

func useDB() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	queries := New(db)

	err = queries.CreateParticipant(context.Background(), CreateParticipantParams{
		Email: "a@a.com",
		Track: "Beginners",
		Name: sql.NullString{
			String: "John",
			Valid:  true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	migrateDB()
	useDB()
}
