package main

import (
	"go-clean-arch/internal/adapter/repository/mock"
	"go-clean-arch/internal/domain/usecase/user"
	"testing"
)

func newTestApplication(t *testing.T) *application {
	t.Helper()

	mockRepo := mock.NewMockRepository()

	userUseCase := user.NewUserService(mockRepo.Users())

	return &application{
		config: config{
			addr: ":8080",
		},
		repository:  mockRepo,
		userUseCase: userUseCase,
	}
}

func getMockRepository(app *application) *mock.MockRepository {
	return app.repository.(*mock.MockRepository)
}
