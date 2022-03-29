# Dependencies

Initiate postgres:

    docker run --rm --name postgres-lesson-7 -e POSTGRES_PASSWORD=postgres -p 5432:5432 postgres

Install migration tool:

    go get -tags 'postgres' -u github.com/golang-migrate/migrate/v4/cmd/migrate/


Run migration via CLI:

    migrate -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path migrations down

Generate Go code:

    sqlc generate
