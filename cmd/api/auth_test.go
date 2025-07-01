package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignIn(t *testing.T) {
	app := newTestApplication(t)
	mux := app.mount()

	t.Run("Should register the user and return a jwt token", func(t *testing.T) {
		userData := map[string]interface{}{
			"name":     "John Doe",
			"email":    "john.doe@email.com",
			"password": "password123",
			"role":     "user",
		}

		jsonData, _ := json.Marshal(userData)

		req, err := http.NewRequest(http.MethodPost, "/v1/auth/sign-in", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("Expected status %d, got %d", http.StatusCreated, rr.Code)
		}

		mockRepo := getMockRepository(app)
		users := mockRepo.MockUsers().GetUsers()

		if len(users) != 1 {
			t.Errorf("expected 1 user, got %d", len(users))
		}

		if users[0].Email != "john.doe@email.com" {
			t.Errorf("Expected email %s, got %s", "john.doe@email.com", users[0].Email)
		}
	})
}
