package user

import "go-clean-arch/internal/adapter/repository"

type UserUseCase interface {
	SigIn(request UserSignInRequest) (*SignInResponse, error)
}

type UserService struct {
	repository repository.UserRepository
}
