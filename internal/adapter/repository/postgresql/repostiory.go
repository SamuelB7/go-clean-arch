package postgresql

import (
	"database/sql"
	"go-clean-arch/internal/adapter/repository"
)

type PostgresUserRepository struct {
	db *sql.DB
}

type PostgresRepository struct {
	userRepo *PostgresUserRepository
}

func NewPostgresRepository(db *sql.DB) repository.Repository {
	return &PostgresRepository{
		userRepo: &PostgresUserRepository{db: db},
	}
}

func (r *PostgresRepository) Users() repository.UserRepository {
	return r.userRepo
}
