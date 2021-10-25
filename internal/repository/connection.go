package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-pg/pg/v10"
)

// Timeout is a Postgres timeout
const Timeout = 5

// DB is a shortcut structure to a Postgres DB
type DB struct {
	*pg.DB
}

// New creates new database connection to postgres
func New(dsn string) (*DB, error) {
	ctx := context.Background()
	pgOpts, err := pg.ParseURL(dsn)
	if err != nil {
		return nil, err
	}
	pgDB := pg.Connect(pgOpts)

	var attempt uint
	const maxAttempts = 10

	for {
		attempt++

		log.Printf("[ping PostgreSQL] (attempt %d)\n", attempt)

		err = pgDB.Ping(ctx)
		if err != nil {
			log.Printf("[ping PostgreSQL] (attempt %d) error: %s\n", attempt, err)

			if attempt < maxAttempts {
				time.Sleep(1 * time.Second)

				continue
			}
			return nil, fmt.Errorf("pgDB.Exec failed: %w", err)
		}
		log.Printf("[PostgreSQL.Dial] (Ping attempt %d) OK\n", attempt)
		break
	}

	pgDB.WithTimeout(time.Second * time.Duration(Timeout))
	return &DB{pgDB}, nil
}
