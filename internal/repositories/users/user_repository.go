package users

import (
	"database/sql"
	"github.com/gattini0928/Equilibrium/internal/models"
)

type UserRepoInterface interface {
	CreateUserWithProfile(user *models.User, patient *models.Patient, therapist *models.Therapist, psychiatrist *models.Psychiatrist) error
	GetUserByEmail(email string) (models.User, error)
}

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}