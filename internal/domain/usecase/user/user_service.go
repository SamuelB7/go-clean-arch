package user

import (
	"context"
	"errors"
	"go-clean-arch/internal/adapter/repository"
	"go-clean-arch/internal/domain/entity"
	"go-clean-arch/internal/domain/usecase/jwt"
)

type UserUseCase interface {
	SignIn(ctx context.Context, request UserSignInRequest) (*SignInResponse, error)
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

func (u *UserService) SignIn(ctx context.Context, request UserSignInRequest) (*SignInResponse, error) {
	if request.Email == "" || request.Password == "" {
		return nil, errors.New("Invalid credentials")
	}

	user, err := u.repository.FindByEmail(request.Email)

	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, errors.New("User already registered")
	}

	user = &entity.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: hashPassword(request.Password),
	}

	if err := u.repository.Create(ctx, user); err != nil {
		return nil, err
	}

	token, err := jwt.GenerateToken(user.Id)

	if err != nil {
		return nil, err
	}

	return &SignInResponse{
		Token: token,
	}, nil
}

func (u *UserService) LogIn(request UserLogInRequest) (*UserLogInResponse, error) {
	panic("unimplemented")
}

func hashPassword(password string) string {
	//TODO: Implement hash function
	return password
}
