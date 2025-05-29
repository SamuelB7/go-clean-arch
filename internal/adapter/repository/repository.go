package repository

import (
	"context"
	"go-clean-arch/internal/domain/entity"
)

type UserRepository interface {
	Create(context.Context, *entity.User) error
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
}

type Repository interface {
	Users() UserRepository
}
