package store

import (
	"context"
	"database/sql"
	"go-clean-arch/internal/domain/entity"
)

type Storage struct {
	Users interface {
		Create(context.Context, *entity.User) error
	}
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Users: &PostgresUserStore{db},
	}
}
