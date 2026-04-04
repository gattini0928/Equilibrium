package users

import (
	"github.com/gattini0928/Equilibrium/internal/models"
)

func (r *UserRepository) GetUserByEmail(email string) (models.User, error) {
	row := r.DB.QueryRow(`
	SELECT user_id, name, email, password, age, cpf, role FROM users WHERE email = $1`, 
	email)
	var user models.User

	err := row.Scan(&user.Name, &user.Email, &user.Password, &user.Age, &user.Cpf, &user.Role)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}