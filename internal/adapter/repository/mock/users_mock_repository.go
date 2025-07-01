package mock

import (
	"context"
	"go-clean-arch/internal/domain/entity"
	"time"
)

type MockUserRepository struct {
	users  []entity.User
	nextID int64
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users:  make([]entity.User, 0),
		nextID: 1,
	}
}

func (m *MockUserRepository) Create(ctx context.Context, user *entity.User) error {
	user.Id = m.nextID
	user.CreatedAt = time.Now().Format("YYYY-MM-DD")
	user.UpdatedAt = time.Now().Format("YYYY-MM-DD")
	m.users = append(m.users, *user)
	m.nextID++
	return nil
}

func (m *MockUserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	for _, user := range m.users {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, nil
}

// Métodos auxiliares para testes
func (m *MockUserRepository) Clear() {
	m.users = make([]entity.User, 0)
	m.nextID = 1
}

func (m *MockUserRepository) GetUsers() []entity.User {
	return m.users
}
