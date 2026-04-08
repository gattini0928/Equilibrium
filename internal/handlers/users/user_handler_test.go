package users

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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
		ID:    1,
		Email: email,
	}, "fake-token-123", nil
}

func (m *MockUserService) CompleteTherapistSignUp(userID int, specialty string, description string) error {
	return nil
}

func (m *MockUserService) CompletePsychiatristSignUp(userID int, specialty string, description string) error {
	return nil
}

func (m *MockUserService) ListAllTherapists() ([]models.DoctorWithUser,error) {
	return []models.DoctorWithUser{}, nil} 

func (m *MockUserService) ListAllPsychiatrists() ([]models.DoctorWithUser,error) {
	return []models.DoctorWithUser{}, nil}
	 

func (m *MockUserService) TherapistDetail(userID int) (models.DoctorWithUser, error) {
	return models.DoctorWithUser{}, nil
}

func (m *MockUserService) PsychiatristDetail(userID int) (models.DoctorWithUser, error) {
	return models.DoctorWithUser{}, nil
}

func (m *MockUserService) TherapistToPatient(patientID int, therapistID int) error {
	return nil
}

func (m *MockUserService) PsychiatristToPatient(patientID int, therapistID int) error {
	return nil
}

func (m *MockUserService) TherapistPatientDetail(userID int) (models.PatientWithUser, error) {
	return models.PatientWithUser{
		ID:               1,
		Name:             "Gabriel Gattini",
		Email:            "gabrielgattini659@gmail.com",
		Age:              29,
		Image:            "profile.png",
		CurrentDiagnosis: "Severe Depression",
		Books: []models.Book{
			{
				ID:     1,
				Author: "Dan Brown",
				Title:  "Ponto de Impacto",
			},
		},
	}, nil
}

func (m *MockUserService) PsychiatristPatientDetail(userID int) (models.PatientWithUser, error) {
	return models.PatientWithUser{
		ID:               1,
		Name:             "Gabriel Gattini",
		Email:            "gabrielgattini659@gmail.com",
		Age:              29,
		Image:            "profile.png",
		CurrentDiagnosis: "Severe Depression",
		Remedies: []models.Remedy{
			{
				ID:     1,
				Name: "Clonazepal",
				Dosage:  "10 gotas dia",
				Quantity: 1,
			},
		},
	}, nil
}

func TestHandleSignUp(t *testing.T) {
	userInput := map[string]any{
		"name":     "Gabriel Gattini",
		"email":    "teste123@gmail.com",
		"password": "Password123$",
		"age":      24,
		"cpf":      "08041289099",
		"role":     "patient",
		"image": "userimage.png",
	}

	body, _ := json.Marshal(userInput)

	req := httptest.NewRequest("POST", "/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handler := NewUserHandler(&MockUserService{})

	mux := http.NewServeMux()
	mux.HandleFunc("POST /signup", handler.HandleSignup)

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", rec.Code)
	}
}

func TestHandleSignUp_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest("POST", "/signup", bytes.NewBuffer([]byte(`{invalid json}`)))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handler := NewUserHandler(&MockUserService{})

	mux := http.NewServeMux()
	mux.HandleFunc("POST /signup", handler.HandleSignup)

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rec.Code)
	}
}

func TestHandleLogin(t *testing.T) {
	userInput := map[string]any{
		"email":    "teste123@gmail.com",
		"password": "Password123$",
	}

	body, _ := json.Marshal(userInput)

	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handler := NewUserHandler(&MockUserService{})

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

func TestHandleLogin_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer([]byte(`{invalid json}`)))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	handler := NewUserHandler(&MockUserService{})

	mux := http.NewServeMux()
	mux.HandleFunc("POST /login", handler.HandleLogin)

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rec.Code)
	}
}