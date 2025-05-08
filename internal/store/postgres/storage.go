package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Users interface {
		Create(context.Context) error
	}
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Users: &PostgresUserStore{db},
	}
}
