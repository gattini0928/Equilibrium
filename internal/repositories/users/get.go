package users

import (
	"github.com/gattini0928/Equilibrium/internal/models"
)

func (r *UserRepository) GetUserByEmail(email string) (models.User, error) {
	row := r.DB.QueryRow(`
		SELECT id, name, email, password, age, cpf, role, image 
		FROM users 
		WHERE email = $1
	`, email)
	var user models.User

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Age,
		&user.Cpf,
		&user.Role,
		&user.Image,
	)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) GetAllTherapists() ([]models.DoctorWithUser, error) {
	query := `
		SELECT 
			t.id,
			u.name,
			u.email,
			u.age,
			u.image,
			t.specialty,
			t.description
		FROM therapists t
		JOIN users u ON t.user_id = u.id; 
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var therapists []models.DoctorWithUser

	for rows.Next() {
		var t models.DoctorWithUser
		err := rows.Scan(
			&t.ID,
			&t.Name,
			&t.Email,
			&t.Age,
			&t.Image,
			&t.Specialty,
			&t.Description,
		)

		if err != nil {
			return nil, err
		}

		therapists = append(therapists, t)
	}

	return therapists, nil
}

func (r *UserRepository) GetAllPsychiatrists() ([]models.DoctorWithUser, error) {
	query := `
		SELECT 
			p.id,
			u.name,
			u.email,
			u.age,
			u.image,
			p.description,
			p.crm
		FROM psychiatrists p
		JOIN users u ON p.user_id = u.id
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var psychiatrists []models.DoctorWithUser

	for rows.Next() {
		var p models.DoctorWithUser
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Email,
			&p.Age,
			&p.Image,
			&p.Description,
			&p.CRM,
		)
		if err != nil {
			return nil, err
		}

		psychiatrists = append(psychiatrists, p)
	}

	return psychiatrists, nil
}