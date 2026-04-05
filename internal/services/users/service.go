package users

import (
    repoUsers "github.com/gattini0928/Equilibrium/internal/repositories/users"
     "github.com/gattini0928/Equilibrium/internal/models"
)

type UserServiceInterface interface {
	CreateUser(user models.User, p models.PatientProfile, t models.TherapistProfile, ps models.PsychiatristProfile) error
	Login(email string, password string) (models.User, string, error)
}

type UserService struct {
    Repo repoUsers.UserRepoInterface
	Secret []byte
}

func NewUserService(r repoUsers.UserRepoInterface, secret []byte) *UserService {
    return &UserService{Repo: r, Secret: secret}
}
