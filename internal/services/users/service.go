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
    CompleteTherapistSignUp(userID int, specialty string, description string) error 
    CompletePsychiatristSignUp(userID int, crm string, description string) error
    TherapistDetail(userID int) (models.DoctorWithUser, error)
    PsychiatristDetail(userID int) (models.DoctorWithUser, error)
    TherapistToPatient(patientID int, therapistID int) error
    PsychiatristToPatient(patientID int, therapistID int) error
}

type UserService struct {
    Repo repoUsers.UserRepoInterface
	Secret []byte
}

func NewUserService(r repoUsers.UserRepoInterface, secret []byte) *UserService {
    return &UserService{Repo: r, Secret: secret}
}
