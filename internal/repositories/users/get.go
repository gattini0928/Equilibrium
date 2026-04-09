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

func (r *UserRepository) GetTherapistById(userID int) (models.DoctorWithUser, error) {
	var therapist models.DoctorWithUser
	query := 
		`
		SELECT 
			t.id,
			u.name,
			u.email,
			u.age,
			u.image,
			t.specialty,
			t.description
		FROM therapists t
		JOIN users u ON t.user_id = u.id
		WHERE t.user_id = $1;	
		`
	
	row := r.DB.QueryRow(query, userID)
		err := row.Scan(
		&therapist.ID,
		&therapist.Name,
		&therapist.Email,
		&therapist.Age,
		&therapist.Image,
		&therapist.Specialty,
		&therapist.Description,
	)

	if err != nil {
		return models.DoctorWithUser{}, err
	}

	return therapist, nil
}

func (r *UserRepository) GetPsychiatristById(userID int) (models.DoctorWithUser, error) {
	var psychiatrist models.DoctorWithUser

	query := `
		SELECT p.id,
			u.name,
			u.email,
			u.age,
			u.image,
			p.crm,
			p.description
		FROM psychiatrists p
		JOIN users u ON p.user_id = u.id
		WHERE p.user_id = $1;
	`
	row := r.DB.QueryRow(query, userID)
	err := row.Scan(
		&psychiatrist.ID,
		&psychiatrist.Name,
		&psychiatrist.Email,
		&psychiatrist.Age,
		&psychiatrist.Image,
		&psychiatrist.CRM,
		&psychiatrist.Description,

	)

	if err != nil {
		return models.DoctorWithUser{}, err
	}

	return psychiatrist, nil
}

func (r *UserRepository) GetTherapistPatientByID(userID int) (models.PatientWithUser, error) {
	var patient models.PatientWithUser
	query := `
	SELECT 
		p.id,
		u.name, u.email, u.age, u.image,
		p.current_diagnosis,
		b.id, b.author, b.title
		FROM patients p
		JOIN users u ON u.id = p.user_id
		JOIN consultations c ON c.patient_id = p.id
		JOIN consultation_books cb ON cb.consultation_id = c.id
		JOIN books b ON b.id = cb.book_id
		WHERE p.user_id = $1;
		`

	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return models.PatientWithUser{}, err
	}

	defer rows.Close()

	var books []models.Book
	for rows.Next(){
		var b models.Book
		err := rows.Scan(
			&patient.ID,
			&patient.Name,
			&patient.Email,
			&patient.Age,
			&patient.Image,
			&patient.CurrentDiagnosis,
			&b.ID,
			&b.Author,
			&b.Title,
		)
		if err != nil {
			return models.PatientWithUser{}, err
		}
		books = append(books, b)
	}

	patient.Books = books
	return patient, nil
}

func (r *UserRepository) GetPsychiatristPatientByID(userID int) (models.PatientWithUser, error) {
	var patient models.PatientWithUser
	query := `
		SELECT p.id, u.name, u.email, u.age, u.image, p.current_diagnosis,
		r.id, r.name, r.dosage, r.quantity
		FROM patients p
		JOIN users u ON u.id = p.id
		JOIN consultation c ON c.patient_id = p.id
		JOIN consultation_remedies cr ON consultation.id = cr.remedy_id
		JOIN remedies r ON r.id = cr.id
		WHERE user_id = $1;
	`

	rows, err := r.DB.Query(query,userID)
	if err != nil {
		return models.PatientWithUser{}, err
	}
	defer rows.Close()

	var remedies []models.Remedy

	for rows.Next() {
		var r models.Remedy
		err := rows.Scan(
			&patient.ID,
			&patient.Name,
			&patient.Email,
			&patient.Age,
			&patient.Image,
			&patient.CurrentDiagnosis,
			&r.ID,
			&r.Name,
			&r.Dosage,
			&r.Quantity,
		)
		if err != nil {
			return models.PatientWithUser{}, err
		}

		remedies = append(remedies, r)
	}

	patient.Remedies = remedies

	return patient, nil
}

func (r *UserRepository) GetAllTherapistPatients(therapist_id int) ([]models.PatientWithUser, error) {
	query := `
	SELECT 
		p.id, u.name, u.email, u.age, u.image, p.current_diagnosis
	FROM patients p
	JOIN users u ON u.id = p.user_id
	WHERE p.therapist_id = (
		SELECT id FROM therapists WHERE user_id = $1
		)
	`

	rows, err := r.DB.Query(query, therapist_id)
	if err != nil {
		return []models.PatientWithUser{}, err
	}
	
	defer rows.Close()

	var patients []models.PatientWithUser

	for rows.Next(){
		var p models.PatientWithUser
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Email,
			&p.Age,
			&p.Image,
			&p.CurrentDiagnosis,
		)
		if err != nil {
			return nil, err
		}

		patients = append(patients, p)
	}

	return patients, nil
}


func (r *UserRepository) GetAllPsichiatristPatients(psychiatrist_id int) ([]models.PatientWithUser, error) {
	query := `
	SELECT 
		p.id, u.name, u.email, u.age, u.image, p.current_diagnosis
	FROM patients p
	JOIN users u ON u.id = p.user_id
	WHERE p.psychiatris_id = (
			SELECT id FROM therapists WHERE user_id = $1
		)
	`

	rows, err := r.DB.Query(query, psychiatrist_id)
	if err != nil {
		return []models.PatientWithUser{}, err
	}
	
	defer rows.Close()

	var patients []models.PatientWithUser

	for rows.Next(){
		var p models.PatientWithUser
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Email,
			&p.Age,
			&p.Image,
			&p.CurrentDiagnosis,
		)
		if err != nil {
			return nil, err
		}

		patients = append(patients, p)
	}

	return patients, nil
}