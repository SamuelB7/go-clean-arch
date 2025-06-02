package postgresql

import (
	"context"
	"database/sql"
	"go-clean-arch/internal/domain/entity"
)

func (s *PostgresUserRepository) Create(ctx context.Context, user *entity.User) error {
	query := `INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id`

	err := s.db.QueryRowContext(ctx, query, user.Name, user.Email, user.Password, user.Role).Scan(&user.Id)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresUserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `SELECT id, name, email, password, role, created_at, updated_at FROM users WHERE email = $1`

	var user entity.User

	err := s.db.QueryRowContext(ctx, query, email).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
