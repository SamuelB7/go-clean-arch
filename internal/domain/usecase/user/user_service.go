package user

import (
	"go-clean-arch/internal/adapter/repository"
)

type UserUseCase interface {
	SigIn(request UserSignInRequest) (*SignInResponse, error)
	LogIn(request UserLogInRequest) (*UserLogInResponse, error)
}

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserUseCase {
	return &UserService{
		repository: repo,
	}
}

func (u *UserService) SigIn(request UserSignInRequest) (*SignInResponse, error) {
	panic("unimplemented")
}

func (u *UserService) LogIn(request UserLogInRequest) (*UserLogInResponse, error) {
	panic("unimplemented")
}
