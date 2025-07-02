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
		}

		jsonData, _ := json.Marshal(userData)

		req, err := http.NewRequest(http.MethodPost, "/v1/auth/sign-in", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		t.Logf("Response status: %d", rr.Code)
		t.Logf("Response body: %s", rr.Body.String())

		if rr.Code != http.StatusCreated {
			t.Errorf("Expected status %d, got %d. Response: %s",
				http.StatusCreated, rr.Code, rr.Body.String())
			return
		}

		mockRepo := getMockRepository(app)
		users := mockRepo.MockUsers().GetUsers()

		if len(users) != 1 {
			t.Errorf("Expected 1 user, got %d", len(users))
			return
		}

		if users[0].Email != "john.doe@email.com" {
			t.Errorf("Expected email %s, got %s", "john.doe@email.com", users[0].Email)
		}
	})

	t.Run("Should return error for invalid request", func(t *testing.T) {
		mockRepo := getMockRepository(app)
		mockRepo.MockUsers().Clear()

		userData := map[string]interface{}{
			"name":     "John Doe",
			"password": "password123",
		}

		jsonData, _ := json.Marshal(userData)
		req, err := http.NewRequest(http.MethodPost, "/v1/auth/sign-in", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		mux.ServeHTTP(rr, req)

		if rr.Code == http.StatusCreated {
			t.Error("Expected error status, got 201 Created")
		}

		users := mockRepo.MockUsers().GetUsers()
		if len(users) != 0 {
			t.Errorf("Expected 0 users after invalid request, got %d", len(users))
		}
	})

	t.Run("Should return error for duplicate email", func(t *testing.T) {

		mockRepo := getMockRepository(app)
		mockRepo.MockUsers().Clear()

		userData1 := map[string]interface{}{
			"name":     "John Doe",
			"email":    "duplicate@email.com",
			"password": "password123",
		}

		jsonData1, _ := json.Marshal(userData1)
		req1, _ := http.NewRequest(http.MethodPost, "/v1/auth/sign-in", bytes.NewBuffer(jsonData1))
		req1.Header.Set("Content-Type", "application/json")

		rr1 := httptest.NewRecorder()
		mux.ServeHTTP(rr1, req1)

		if rr1.Code != http.StatusCreated {
			t.Logf("First user creation failed: %s", rr1.Body.String())
		}

		userData2 := map[string]interface{}{
			"name":     "Jane Doe",
			"email":    "duplicate@email.com",
			"password": "password456",
		}

		jsonData2, _ := json.Marshal(userData2)
		req2, _ := http.NewRequest(http.MethodPost, "/v1/auth/sign-in", bytes.NewBuffer(jsonData2))
		req2.Header.Set("Content-Type", "application/json")

		rr2 := httptest.NewRecorder()
		mux.ServeHTTP(rr2, req2)

		if rr2.Code == http.StatusCreated {
			t.Error("Expected error for duplicate email, got 201 Created")
		}

		users := mockRepo.MockUsers().GetUsers()
		if len(users) != 1 {
			t.Errorf("Expected 1 user after duplicate email test, got %d", len(users))
		}
	})
}
