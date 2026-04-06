package users

import (
    repoUsers "github.com/gattini0928/Equilibrium/internal/repositories/users"
     "github.com/gattini0928/Equilibrium/internal/models"
)

type UserServiceInterface interface {
	CreateUser(user models.User, p models.Patient, t models.Therapist, ps models.Psychiatrist) error
	Login(email string, password string) (models.User, string, error)
    ListAllTherapists() ([]models.DoctorWithUser, error)
    ListAllPsychiatrists() ([]models.DoctorWithUser, error)
}

type UserService struct {
    Repo repoUsers.UserRepoInterface
	Secret []byte
}

func NewUserService(r repoUsers.UserRepoInterface, secret []byte) *UserService {
    return &UserService{Repo: r, Secret: secret}
}
