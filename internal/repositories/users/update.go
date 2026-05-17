package users

import (
	"database/sql"
	"time"
	"errors"
)

func (r *UserRepository) CompleteTherapist(userID int, specialty string, description string, price float64) error {
	_, err := r.DB.Exec(`
		UPDATE therapists
		SET specialty = $1, description = $2, price = $3
		WHERE user_id = $4;
	`, specialty, description, price, userID)

	return err
}

func (r *UserRepository) CompletePsychiatrist(userID int, crm string, description string, price float64) error {
	_, err := r.DB.Exec(`
		UPDATE psychiatrists
		SET crm = $1, description = $2, price = $3
		WHERE user_id = $3;
	`, crm, description, price, userID)

	return err
}

func (r *UserRepository) AddTherapistToPatient(patientID int, therapistID int) error {
	_, err := r.DB.Exec(`
		UPDATE patients
		SET therapist_id = $1
		WHERE user_id = $2;
	`, therapistID, patientID)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) AddPsychiatristToPatient(patientID int, psychiatristID int) error {
	_, err := r.DB.Exec(`
		UPDATE patients
		SET psychiatrist_id = $1
		WHERE user_id = $2
	`, psychiatristID, patientID)

	return err
}

func (r *UserRepository) UpdateTherapistPrice(userID int, price float64) error {
	_, err := r.DB.Exec(`
		UPDATE therapists 
		SET price = $1 
		WHERE user_id = $2
	`, price, userID)
	return err
}

func (r *UserRepository) UpdatePsychiatristPrice(userID int, price float64) error {
	_, err := r.DB.Exec(`
		UPDATE psychiatrists 
		SET price = $1 
		WHERE user_id = $2
	`, price, userID)
	return err
}

func (r *UserRepository) MarkAgendaReserved(tx *sql.Tx, agendaID int, patientID int) error {
	res, err := tx.Exec(`
		UPDATE agendas
		SET reserved = true,
		    patient_id = $2
		WHERE id = $1
		AND reserved = false
	`, agendaID, patientID)
	
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("agenda já reservada")
	}

	return nil
}

func (r *UserRepository) UpdateConsultationInProgress(consultationID int) error {
	_, err := r.DB.Exec(`
		UPDATE consultations
		SET status = 'in_progress'
		WHERE id = $1
	
	`, consultationID)

	return err
}

func (r *UserRepository) UpdateConsultationFinished(consultationID int) error {
	_, err := r.DB.Exec(`
		UPDATE consultations
		SET status = 'finished',
			date = $1
		WHERE id = $2
	
	`, time.Now(), consultationID)

	return err
}

func (r *UserRepository) UpdateAnnotationConsultation(consultationID int, annotation string) error {
	_, err := r.DB.Exec(`
		UPDATE consultations
		SET annotation = $1
		WHERE id = $2
	`, annotation, consultationID)

	return err
}