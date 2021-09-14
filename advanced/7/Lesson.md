# Dependencies

Initiate postgres:

    docker run --rm --name postgres-lesson-5 -e POSTGRES_PASSWORD=postgres -p 5432:5432 postgres

Install migration tool:

    go get -tags 'postgres' -u github.com/golang-migrate/migrate/v4/cmd/migrate/


Run migration via CLI:

    migrate -database postgres://postgres:postgres@localhost:5432/postgres -path ./migrations up

Generate Go code:

    sqlc generate
