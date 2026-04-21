package users

import "errors"

// "github.com/gattini0928/Equilibrium/internal/models"

func (r *UserRepository) DeleteAgenda(userID int, agendaID int) error {
	query := `
		DELETE FROM agendas 
		WHERE id = $1
		AND professional_id = $2
	`
	res, err := r.DB.Exec(query, agendaID, userID)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return errors.New("agenda não encontrada ou não pertence ao usuário")
	}

	return nil
}
