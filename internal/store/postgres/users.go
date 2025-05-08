package store

import (
	"context"
	"database/sql"
)

type PostgresUserStore struct {
	db *sql.DB
}

func (s *PostgresUserStore) Create(ctx context.Context) error {
	return nil
}
