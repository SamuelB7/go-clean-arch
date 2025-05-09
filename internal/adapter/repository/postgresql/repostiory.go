package postgresql

import (
	"context"
	"database/sql"
	"go-clean-arch/internal/domain/entity"
)

type Repository struct {
	Users interface {
		Create(context.Context, *entity.User) error
	}
}

func NewPostgresRepository(db *sql.DB) Repository {
	return Repository{
		Users: &PostgresUserRepository{db},
	}
}
