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
