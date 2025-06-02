package user

import (
	"context"
	"errors"
	"go-clean-arch/internal/adapter/repository"
	"go-clean-arch/internal/domain/entity"
	"go-clean-arch/internal/domain/usecase/jwt"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	SignIn(ctx context.Context, request UserSignInRequest) (*SignInResponse, error)
	LogIn(ctx context.Context, request UserLogInRequest) (*UserLogInResponse, error)
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

	user, err := u.repository.FindByEmail(ctx, request.Email)

	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, errors.New("User already registered")
	}

	hashedPAssword, err := hashPassword(request.Password)

	if err != nil {
		return nil, errors.New("Error processing password")
	}

	user = &entity.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: hashedPAssword,
		Role:     "CLIENT",
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

func (u *UserService) LogIn(ctx context.Context, request UserLogInRequest) (*UserLogInResponse, error) {
	if request.Email == "" || request.Password == "" {
		return nil, errors.New("Invalid credentials")
	}

	user, err := u.repository.FindByEmail(ctx, request.Email)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("User not found")
	}

	if !verifyPassword(user.Password, request.Password) {
		return nil, errors.New("Invalid credentials")
	}

	token, err := jwt.GenerateToken(user.Id)

	if err != nil {
		return nil, err
	}

	return &UserLogInResponse{
		Token: token,
	}, nil
}

func hashPassword(password string) (string, error) {
	const cost = 12

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func verifyPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
