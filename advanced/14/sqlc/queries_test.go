package sqlc

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/ory/dockertest/v3"
)

func NewTestDB(t *testing.T) *Queries {
	p, err := dockertest.NewPool("")
	if err != nil {
		t.Fatalf("failed to setup pool: %v\n", err)
	}

	res, err := p.Run("postgres", "14-alpine", []string{"POSTGRES_PASSWORD=postgres"})
	if err != nil {
		t.Fatalf("failed to container pool: %v\n", err)
	}

	var db *sql.DB
	if err := p.Retry(func() error {
		var err error
		db, err = sql.Open(
			"postgres",
			fmt.Sprintf("postgres://postgres:postgres@localhost:%s/postgres?sslmode=disable", res.GetPort("5432/tcp")),
		)
		if err != nil {
			return err
		}

		return db.Ping()
	}); err != nil {
		t.Fatalf("failed to connect to container: %v\n", err)
	}

	t.Cleanup(func() {
		if err := p.Purge(res); err != nil {
			t.Fatalf("failed to purge to container: %v\n", err)
		}
	})

	return &Queries{db: db}
}

func TestCreateParticipant(t *testing.T) {

	t.Run("test", func(t *testing.T) {
		db.Refresh()
		NewTestDB(t)
	})
}
