package users

import (
	"errors"

	"github.com/gattini0928/Equilibrium/internal/models"
)

func (r *UserRepository) CreateUserWithProfile(
	user *models.User,
	patient *models.Patient,
	therapist *models.Therapist,
	psychiatrist *models.Psychiatrist,
) error {

	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	var userID int

	err = tx.QueryRow(`
		INSERT INTO users (name, email, password, age, cpf, role, image)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id;
	`, user.Name, user.Email, user.Password, user.Age, user.Cpf, user.Role, user.Image).Scan(&userID)

	if err != nil {
		tx.Rollback()
		return err
	}

	switch user.Role {

	case "patient":
		_, err = tx.Exec(`
			INSERT INTO patients (user_id)
			VALUES ($1);
		`, userID)

	case "therapist":
		_, err = tx.Exec(`
			INSERT INTO therapists (user_id)
			VALUES ($1);
		`, userID)

	case "psychiatrist":
		_, err = tx.Exec(`
			INSERT INTO psychiatrists (user_id)
			VALUES ($1);
		`, userID)

	default:
		tx.Rollback()
		return errors.New("invalid role")
	}

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *UserRepository) CompleteTherapist(userID int, specialty string, description string) error {
	_, err := r.DB.Exec(`
		UPDATE therapists
		SET specialty = $1, description = $2
		WHERE user_id = $3;
	`, specialty, description, userID)

	return err
}

func (r *UserRepository) CompletePsychiatrist(userID int, crm string, description string) error {
	_, err := r.DB.Exec(`
		UPDATE psychiatrists
		SET crm = $1, description = $2
		WHERE user_id = $3;
	`, crm, description, userID)

	return err
}
