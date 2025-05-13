package postgresql

import (
	"context"
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
