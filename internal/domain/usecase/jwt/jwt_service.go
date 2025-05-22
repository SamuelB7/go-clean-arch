package jwt

import "go-clean-arch/internal/domain/entity"

type JwtService interface {
	GenerateToken(user *entity.User) (string error)
}
