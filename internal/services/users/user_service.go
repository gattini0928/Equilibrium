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

func (s *UserService) CompleteTherapistSignUp(userID int, specialty string, description string) error {
	return s.Repo.CompleteTherapist(userID, specialty, description)
}

func (s *UserService) CompletePsychiatristSignUp(userID int, crm string, description string) error {
	return s.Repo.CompletePsychiatrist(userID, crm, description)
}

func (s *UserService) ListAllTherapists() ([]models.DoctorWithUser, error) {
	therapists, err := s.Repo.GetAllTherapists()
	if err != nil {
		return nil, err
	}

	return therapists, nil
}


func (s *UserService) ListAllPsychiatrists() ([]models.DoctorWithUser, error) {
	psychiatrists, err := s.Repo.GetAllPsychiatrists()
	if err != nil {
		return nil, err
	}
	
	return psychiatrists, nil
}

func (s *UserService) TherapistDetail(userID int) (models.DoctorWithUser, error) {
	therapist, err := s.Repo.GetTherapistById(userID)
	if err != nil {
		return models.DoctorWithUser{}, err
	}

	return therapist, nil

}

func (s *UserService) PsychiatristDetail(userID int) (models.DoctorWithUser, error) {
	psychiatrist, err := s.Repo.GetPsychiatristById(userID)
	if err != nil {
		return models.DoctorWithUser{}, err
	}

	return psychiatrist, nil

}
