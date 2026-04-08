package users

import (
	"testing"

	"github.com/gattini0928/Equilibrium/internal/models"
	validators "github.com/gattini0928/Equilibrium/internal/services/validators"

)

type MockRepository struct {}

func (m *MockRepository) CreateUserWithProfile(	
	user *models.User,
	patient *models.Patient,
	therapist *models.Therapist,
	psychiatrist *models.Psychiatrist) error {
		return nil
	}

func (m *MockRepository) CompleteTherapist(userID int, crm string, description string) error {
	return nil
}

func (m *MockRepository) CompletePsychiatrist(userID int, specialty string, description string) error {
	return nil
}

func (m *MockRepository) GetTherapistById(userID int) (models.DoctorWithUser, error){
	return models.DoctorWithUser{}, nil
}

func (m *MockRepository) GetPsychiatristById(userID int) (models.DoctorWithUser, error){
	return models.DoctorWithUser{}, nil
}

func (m *MockRepository) AddTherapistToPatient(patientID int, therapistID int) error {
	return nil
}

func (m *MockRepository) AddPsychiatristToPatient(patientID int, therapistID int) error {
	return nil
}

var hash, _ = validators.HashPassword("Password123$")

func (m *MockRepository) GetUserByEmail(email string) (models.User, error) {
	return models.User{
		ID:1,
		Email:email,
		Password: hash,
	}, nil}

func (m *MockRepository) GetAllTherapists() ([]models.DoctorWithUser, error) {
	return []models.DoctorWithUser{}, nil
}

func (m *MockRepository) GetAllPsychiatrists() ([]models.DoctorWithUser, error) {
	return []models.DoctorWithUser{}, nil
}

func TestCreateUser_Success(t *testing.T) {
	service := NewUserService(&MockRepository{}, []byte("secret"))

	user := models.User{
		Name:     "Gabriel Gattini",
		Email:    "teste123@gmail.com",
		Password: "Password123$",
		Age:      24,
		Cpf:      "12312441309",
		Role:     "patient",
	}

	var patient models.Patient
	var therapist models.Therapist
	var psychiatrist models.Psychiatrist

	err := service.CreateUser(user, patient, therapist, psychiatrist)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestCreateUser_InvalidName(t *testing.T) {
	service := NewUserService(&MockRepository{}, []byte("secret"))

	user := models.User{
		Name:     "Gabriel",
		Email:    "teste123@gmail.com",
		Password: "Password123$",
		Age:      24,
		Cpf:      "12312441309",
		Role:     "patient",
	}

	err := service.CreateUser(user, models.Patient{}, models.Therapist{}, models.Psychiatrist{})

	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestCreateUser_InvalidEmail(t *testing.T) {
	service := NewUserService(&MockRepository{}, []byte("secret"))

	user := models.User{
		Name:     "Gabriel Gattini",
		Email:    "email-errado",
		Password: "Password123$",
		Age:      24,
		Cpf:      "12312441309",
		Role:     "patient",
	}

	err := service.CreateUser(user, models.Patient{}, models.Therapist{}, models.Psychiatrist{})

	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestCreateUser_InvalidPassword(t *testing.T) {
	service := NewUserService(&MockRepository{}, []byte("secret"))

	user := models.User{
		Name:     "Gabriel Gattini",
		Email:    "teste123@gmail.com",
		Password: "Password123",
		Age:      24,
		Cpf:      "12312441309",
		Role:     "patient",
	}

	err := service.CreateUser(user, models.Patient{}, models.Therapist{}, models.Psychiatrist{})

	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestCreateUser_InvalidCpf(t *testing.T) {
	service := NewUserService(&MockRepository{}, []byte("secret"))

	user := models.User{
		Name:     "Gabriel Gattini",
		Email:    "teste123@gmail.com",
		Password: "Password123$",
		Age:      24,
		Cpf:      "1231244130",
		Role:     "patient",
	}

	err := service.CreateUser(user, models.Patient{}, models.Therapist{}, models.Psychiatrist{})

	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestCreateUser_InvalidAge(t *testing.T) {
	service := NewUserService(&MockRepository{}, []byte("secret"))

	user := models.User{
		Name:     "Gabriel Gattini",
		Email:    "teste123@gmail.com",
		Password: "Password123$",
		Age:      18,
		Cpf:      "12312441309",
		Role:     "therapist",
		Image: 		"userimage.png",
	}

	err := service.CreateUser(user, models.Patient{}, models.Therapist{}, models.Psychiatrist{})

	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestLogin(t *testing.T) {
	service := NewUserService(&MockRepository{}, []byte("secret"))

	email := "teste123@gmail.com"
	password := "Password123$"

	user, token, err := service.Login(email, password)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if user.Email != email {
		t.Errorf("expected email %s, got %s", email, user.Email)
	}

	if token == "" {
		t.Errorf("expected token, got empty string")
	}
}

func TestLogin_InvalidPassword(t *testing.T) {
	service := NewUserService(&MockRepository{}, []byte("secret"))

	email := "teste123@gmail.com"
	password := "senhaErrada"

	_, token, err := service.Login(email, password)

	if err == nil {
		t.Errorf("expected error, got nil")
	}

	if token != "" {
		t.Errorf("expected empty token, got %s", token)
	}
}