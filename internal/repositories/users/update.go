package users

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

func (r *UserRepository) AddTherapistToPatient(patientID int, therapistID int) error {
	_, err := r.DB.Exec(`
		UPDATE public.patients 
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