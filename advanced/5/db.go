package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func connect() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func createTable(db *sql.DB) {
	stm := `CREATE TABLE "user" (
				email text,
				track text
			)`

	_, err := db.Exec(stm)
	if err != nil {
		log.Fatal(err)
	}
}

func saveUser(db *sql.DB, email, track string) {
	stm := `INSERT INTO "user" VALUES ($1, $2)`

	_, err := db.Exec(stm, email, track)
	if err != nil {
		log.Fatal(err)
	}
}

func getUserTrack(db *sql.DB, email string) string {
	query := `SELECT track FROM "user" WHERE email = $1`

	row := db.QueryRow(query, email)
	var track string
	err := row.Scan(&track)
	if err != nil {
		log.Fatal(err)
	}
	return track
}

func dropTable(db *sql.DB) {
	_, err := db.Exec(`DROP TABLE "user"`)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.SetFlags(log.Llongfile)
	db := connect()
	dropTable(db)
	createTable(db)
	saveUser(db, "frodo@hobbits.com", "Advanced")
	track := getUserTrack(db, "frodo@hobbits.com")

	fmt.Println("got track: ", track)
}
