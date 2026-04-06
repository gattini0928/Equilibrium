package routes

import (
	"strings"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	userHandlerPkg  "github.com/gattini0928/Equilibrium/internal/handlers/users"
	"github.com/gattini0928/Equilibrium/internal/models"

)

type MockUserService struct{}

func (m *MockUserService) CreateUser(
	user models.User,
	p models.Patient,
	t models.Therapist,
	ps models.Psychiatrist,
) error {
	return nil
}

func (m *MockUserService) Login(email string, password string) (models.User, string, error) {
	return models.User{
		ID: 1,
		Email: email,
	}, "fake-token-123", nil
}

func TestSignUpRoute(t *testing.T) {
	userInput := map[string]any{
		"name":     "Gabriel Gattini",
		"email":    "teste123@gmail.com",
		"password": "Password123$",
		"age":      24,
		"cpf":      "09045108010",
		"role":     "patient",
		"image": "userimg.png",
	}

	body, _ := json.Marshal(userInput)

	req := httptest.NewRequest("POST", "/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	mockService := &MockUserService{}

	handler := userHandlerPkg.NewUserHandler(mockService)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /signup", handler.HandleSignup)

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", rec.Code)
	}
}

func TestLoginRoute(t *testing.T) {
	userInput := map[string]any{
		"email":    "teste123@gmail.com",
		"password": "Password123$",
	}

	body, _ := json.Marshal(userInput)

	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	mockService := &MockUserService{}

	handler := userHandlerPkg.NewUserHandler(mockService)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /login", handler.HandleLogin)

	mux.ServeHTTP(rec, req)


	if rec.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rec.Code)
	}

	respBody := rec.Body.String()

	if !strings.Contains(respBody, "fake-token-123") {
		t.Errorf("expected token in response, got %s", respBody)
	}
}