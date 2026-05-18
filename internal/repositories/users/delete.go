package users

import (
	"database/sql"
	"errors"
)



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

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	
	if rows == 0 {
		return errors.New("agenda não encontrada ou não pertence ao usuário")
	}

	return nil
}

func (r *UserRepository) DeleteAgendaConsultation(tx *sql.Tx ,userID int, agendaID int) error {
	query := `
		DELETE FROM agendas 
		WHERE id = $1
		AND professional_id = $2
	`
	res, err := r.DB.Exec(query, agendaID, userID)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	
	if rows == 0 {
		return errors.New("agenda não encontrada ou não pertence ao usuário")
	}

	return nil
}


