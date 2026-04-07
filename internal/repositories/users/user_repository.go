package users

import (
	"database/sql"
	"github.com/gattini0928/Equilibrium/internal/models"
)

type UserRepoInterface interface {
	CreateUserWithProfile(user *models.User, patient *models.Patient, therapist *models.Therapist, psychiatrist *models.Psychiatrist) error
	GetUserByEmail(email string) (models.User, error)
	GetAllTherapists() ([]models.DoctorWithUser, error)
	GetAllPsychiatrists() ([]models.DoctorWithUser, error)
	CompleteTherapist(userID int, specialty string, description string) error
	CompletePsychiatrist(userID int, crm string, description string) error
	GetTherapistById(userID int) (models.DoctorWithUser, error)
	GetPsychiatristById(userID int) (models.DoctorWithUser, error)
}

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}