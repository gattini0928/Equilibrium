package users

import (
	"errors"

	"github.com/gattini0928/Equilibrium/internal/configs"
	"github.com/gattini0928/Equilibrium/internal/models"
	"github.com/gattini0928/Equilibrium/internal/services/auth"
	"github.com/gattini0928/Equilibrium/internal/services/validators"
)

var (
	ErrInvalidInput = errors.New("invalid input")
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidPassword = errors.New("invalid password")
	ErrTokenFailed = errors.New("token failed")
)

func (u *UserService) CreateUser(user models.User, patient models.Patient, therapist models.Therapist, psychiatrist models.Psychiatrist) error {
	var err error

	err = validators.ValidateName(user.Name)
	if err != nil {
		return err
	}
	err = validators.ValidateEmail(user.Email)
	if err != nil {
		return err
	}

	err = validators.ValidatePassword(user.Password)
	if err != nil {
		return err
	}

	if user.Role == "therapist" || user.Role == "psychiatrist" {
		err = validators.ValidateAge(user.Age, user.Role)
		if err != nil {
			return err
		}
	}

	err = validators.ValidateCpf(user.Cpf)
	if err != nil {
		return err
	}

	hashPassword, err := validators.HashPassword(user.Password)
	if err != nil {
		return ErrInvalidPassword
	}

	user.Password = hashPassword

	return u.Repo.CreateUserWithProfile(&user, &patient, &therapist, &psychiatrist)
}

func (s *UserService) Login(email string, password string) (models.User, string, error) {
	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		return models.User{}, "", err
	}

	if !validators.CheckPasswordHash(password, user.Password){
		return models.User{}, "", ErrInvalidPassword
	}

	cfg := configs.LoadDBConfig()

	token, err := auth.CreateJWT(s.Secret, user.ID, cfg.JWTExpirationInSeconds)
	if err != nil {
		return models.User{}, "", ErrTokenFailed
	}

	return user, token, nil
}

// Completar cadastro do therapeuta
func (s *UserService) CompleteTherapistSignUp(userID int, specialty string, description string) error {
	return s.Repo.CompleteTherapist(userID, specialty, description)
}

// Completar cadastro do psiquiatra
func (s *UserService) CompletePsychiatristSignUp(userID int, crm string, description string) error {
	return s.Repo.CompletePsychiatrist(userID, crm, description)
}

// Listagem de todos terapeutas
func (s *UserService) ListAllTherapists() ([]models.DoctorWithUser, error) {
	return s.Repo.GetAllTherapists()
}
// Listagem de todos psiquiatras
func (s *UserService) ListAllPsychiatrists() ([]models.DoctorWithUser, error) {
	return s.Repo.GetAllPsychiatrists()
}

// Detalhes do terapeuta
func (s *UserService) TherapistDetail(userID int) (models.DoctorWithUser, []models.Agenda, error) {

	therapist, err := s.Repo.GetTherapistById(userID)
	if err != nil {
		return models.DoctorWithUser{}, nil, err
	}

	agendas, err := s.Repo.GetTherapistAgenda(therapist.ID)
	if err != nil {
		return models.DoctorWithUser{}, nil, err
	}

	return therapist, agendas, nil
}

// Detalhes do psiquiatra
func (s *UserService) PsychiatristDetail(userID int) (models.DoctorWithUser, []models.Agenda, error) {

	psychiatrist, err := s.Repo.GetPsychiatristById(userID)
	if err != nil {
		return models.DoctorWithUser{}, nil, err
	}

	agendas, err := s.Repo.GetTherapistAgenda(psychiatrist.ID)
	if err != nil {
		return models.DoctorWithUser{}, nil, err
	}

	return psychiatrist, agendas, nil
}

func (s *UserService) TherapistToPatient(patientID, therapistID int) error {
	user, err := s.Repo.GetUserByID(patientID)

	if err != nil {
		return err
	}

	if user.Role != "patient" {
		return errors.New("forbidden")
	}
	
	return s.Repo.AddTherapistToPatient(patientID, therapistID)
}

func (s *UserService) PsychiatristToPatient(patientID, psychiatristID int) error {
	user, err := s.Repo.GetUserByID(patientID)
	if err != nil {
		return err
	}

	if user.Role != "patient" {
		return errors.New("forbidden")
	}

	return s.Repo.AddPsychiatristToPatient(patientID, psychiatristID)

}

// Terapeuta ou Psiquiatra vê os detalhes do paciente
func (s *UserService) GetPatientDetail(patientID, doctorID int) (models.PatientWithUser, error) {
	user, err := s.Repo.GetUserByID(doctorID)
	if err != nil {
		return models.PatientWithUser{}, err
	}

	if user.Role == "therapist" {
		return s.Repo.GetTherapistPatient(patientID)
	}

	if user.Role == "psychiatrist" {
		return s.Repo.GetPsychiatristPatient(patientID)
	}

	return models.PatientWithUser{}, errors.New("forbidden")
}	

// Detalhes do Terapeuta do Paciente
func (s *UserService) PatientTherapistDetail(userID int) (models.DoctorWithUser, error) {
	return s.Repo.GetPatientTherapist(userID)
}

// Detalhes do Psiquiatra do Paciente
func (s *UserService) PatientPsiquiatristDetail(userID int) (models.DoctorWithUser, error) {
	return s.Repo.GetPatientPsychiatrist(userID) 
}

// Listar pacientes do terapeuta ou psiquiatra
func (s *UserService) ListMyPatients(userID int) ([]models.PatientWithUser, error) {
	user, err := s.Repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	switch user.Role {
	case "therapist":
		return s.Repo.GetTherapistPatients(userID)

	case "psychiatrist":
		return s.Repo.GetPsychiatristPatients(userID)

	default:
		return nil, errors.New("forbidden")
	}
}

func (s *UserService) AddAgenda(userID int, day int, month int, hour string) (models.Agenda, error) {
	user, err := s.Repo.GetUserByID(userID)
	if err != nil {
		return models.Agenda{}, err
	}

	var professionalID int

	switch user.Role {

	case "therapist":
		professionalID, err = s.Repo.GetTherapistIDByUserID(userID)

	case "psychiatrist":
		professionalID, err = s.Repo.GetPsychiatristIDByUserID(userID)

	default:
		return models.Agenda{}, errors.New("forbidden")
	}

	if err != nil {
		return models.Agenda{}, err
	}

	return s.Repo.InsertAgenda(professionalID, day, month, hour)
}

