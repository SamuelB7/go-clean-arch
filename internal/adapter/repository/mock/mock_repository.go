package mock

import "go-clean-arch/internal/adapter/repository"

type MockRepository struct {
	userRepo *MockUserRepository
}

func NewMockRepository() repository.Repository {
	return &MockRepository{
		userRepo: NewMockUserRepository(),
	}
}

func (m *MockRepository) Users() repository.UserRepository {
	return m.userRepo
}

func (m *MockRepository) MockUsers() *MockUserRepository {
	return m.userRepo
}
