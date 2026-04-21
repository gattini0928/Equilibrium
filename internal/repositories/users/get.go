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

func (r *UserRepository) GetUserByID(id int) (models.User, error) {
	var user models.User

	row := r.DB.QueryRow(`
		SELECT id, name, email, age, image, role
		FROM users 
		WHERE id = $1`, id)

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Age,
		&user.Image,
		&user.Role,
	)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) GetPatientIDByUserID(userID int) (int, error) {
	var patientID int

	err := r.DB.QueryRow(`
		SELECT id
		FROM patients
		WHERE user_id = $1	
	`, userID).Scan(&patientID)

	if err != nil {
		return 0, err
	}

	return patientID, nil
}

func (r *UserRepository) GetPatientPerfil(userID int) (models.UserPerfil, error) {
	query := `
		SELECT u.id, u.name, u.email, u.age, u.image, p.current_diagnosis
		FROM patients p
		JOIN users u ON p.user_id = u.id
		WHERE p.user_id = $1
	`
	var patient models.UserPerfil

	row := r.DB.QueryRow(query, userID)
	err := row.Scan(
		&patient.ID,
		&patient.Name,
		&patient.Email,
		&patient.Age,
		&patient.Image,
		&patient.CurrentDiagnosis,
	)

	if err != nil {
		return models.UserPerfil{}, err
	}

	return patient, nil
}

func (r *UserRepository) GetTherapistPerfil(userID int) (models.UserPerfil, error) {
	query := `
		SELECT u.id, u.name, u.email, u.age, u.image, u.role, t.specialty, t.description, t.price
		FROM therapists t
		JOIN users u ON t.user_id = u.id
		WHERE t.user_id = $1
	`
	var therapist models.UserPerfil

	row := r.DB.QueryRow(query, userID)
	err := row.Scan(
		&therapist.ID,
		&therapist.Name,
		&therapist.Email,
		&therapist.Age,
		&therapist.Image,
		&therapist.Role,
		&therapist.Specialty,
		&therapist.Description,
		&therapist.Price,
	)

	if err != nil {
		return models.UserPerfil{}, err
	}

	return therapist, nil
}

func (r *UserRepository) GetPsychiatristPerfil(userID int) (models.UserPerfil, error) {
	query := `
		SELECT u.id, u.name, u.email, u.age, u.image, u.role, p.crm, p.description, p.price
		FROM psychiatrists p
		JOIN users u ON p.user_id = u.id
		WHERE p.user_id = $1
	`
	var psychiatrist models.UserPerfil

	row := r.DB.QueryRow(query, userID)
	err := row.Scan(
		&psychiatrist.ID,
		&psychiatrist.Name,
		&psychiatrist.Email,
		&psychiatrist.Age,
		&psychiatrist.Image,
		&psychiatrist.Role,
		&psychiatrist.CRM,
		&psychiatrist.Description,
		&psychiatrist.Price,
	)

	if err != nil {
		return models.UserPerfil{}, err
	}

	return psychiatrist, nil
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
		WHERE t.description IS NOT NULL AND t.crm IS NOT NULL 
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
		WHERE p.description IS NOT NULL AND p.crm IS NOT NULL
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

func (r *UserRepository) GetTherapistAgenda(therapistID int) ([]models.Agenda, error ) {
	query := `
		SELECT day, month, hour, reserved
		FROM agendas
		WHERE professional_id = $1
		AND reserved = false
	`
	rows, err := r.DB.Query(query,therapistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var agendas []models.Agenda

	for rows.Next() {
		var a models.Agenda
		err := rows.Scan(
			&a.Day,
			&a.Month,
			&a.Hour,
			&a.Reserved,
		)
		if err != nil {
			return nil, err
		}
		agendas = append(agendas, a)
	}

	return agendas, nil
}


func (r *UserRepository) GetTherapistPrivateAgenda(userID int) ([]models.Agenda, error ) {
	query := `
		SELECT day, month, hour, reserved
		FROM agendas
		WHERE professional_id = $1
	`
	rows, err := r.DB.Query(query,userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var agendas []models.Agenda

	for rows.Next() {
		var a models.Agenda
		err := rows.Scan(
			&a.Day,
			&a.Month,
			&a.Hour,
			&a.Reserved,
		)
		if err != nil {
			return nil, err
		}
		agendas = append(agendas, a)
	}

	return agendas, nil
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

func (r *UserRepository) GetPsychiatristAgenda(psychiatristID int) ([]models.Agenda, error ) {
	query := `
		SELECT day, month, hour, reserved
		FROM agendas
		WHERE professional_id = $1
		AND reserved = false
	`
	rows, err := r.DB.Query(query,psychiatristID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var agendas []models.Agenda

	for rows.Next() {
		var a models.Agenda
		err := rows.Scan(
			&a.Day,
			&a.Month,
			&a.Hour,
			&a.Reserved,
		)
		if err != nil {
			return nil, err
		}
		agendas = append(agendas, a)
	}

	return agendas, nil
}

func (r *UserRepository) GetPsychiatristPrivateAgenda(userID int) ([]models.Agenda, error ) {
	query := `
		SELECT day, month, hour, reserved
		FROM agendas
		WHERE professional_id = $1
	`
	rows, err := r.DB.Query(query,userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var agendas []models.Agenda

	for rows.Next() {
		var a models.Agenda
		err := rows.Scan(
			&a.Day,
			&a.Month,
			&a.Hour,
			&a.Reserved,
		)
		if err != nil {
			return nil, err
		}
		agendas = append(agendas, a)
	}

	return agendas, nil
}

func (r *UserRepository) GetTherapistPatient(patiendID int) (models.PatientWithUser, error) {
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

	rows, err := r.DB.Query(query, patiendID)
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

func (r *UserRepository) GetPsychiatristPatient(patiendID int) (models.PatientWithUser, error) {
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

	rows, err := r.DB.Query(query,patiendID)
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

 func (r *UserRepository) GetTherapistPatients(doctorID int) ([]models.PatientWithUser, error) {
	query := `
	SELECT p.id, u.name, u.email, u.age, u.image, p.current_diagnosis
	FROM patients p
	JOIN users u ON u.id = p.user_id
	JOIN therapists t ON t.id = p.therapist_id
	WHERE t.user_id = $1
	`

	rows, err := r.DB.Query(query, doctorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []models.PatientWithUser

	for rows.Next() {
		var p models.PatientWithUser
		err := rows.Scan(&p.ID, &p.Name, &p.Email, &p.Age, &p.Image, &p.CurrentDiagnosis)
		if err != nil {
			return nil, err
		}
		patients = append(patients, p)
	}

	return patients, nil
}

 func (r *UserRepository) GetPsychiatristPatients(doctorID int) ([]models.PatientWithUser, error) {
	query := `
	SELECT p.id, u.name, u.email, u.age, u.image, p.current_diagnosis
	FROM patients p
	JOIN users u ON u.id = p.user_id
	JOIN psychiatrists ps ON t.id = p.psychiastrid_id
	WHERE ps.user_id = $1
	`

	rows, err := r.DB.Query(query, doctorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []models.PatientWithUser

	for rows.Next() {
		var p models.PatientWithUser
		err := rows.Scan(
			&p.ID, &p.Name, &p.Email, &p.Age, &p.Image, 
			&p.CurrentDiagnosis)

		if err != nil {
			return nil, err
		}

		patients = append(patients, p)
	}

	return patients, nil
}

func (r *UserRepository) GetPatientTherapist(userID int) (models.DoctorWithUser, error) {
	var therapist models.DoctorWithUser

	query := `
	SELECT 
		t.id,
		u.name,
		u.email,
		u.age,
		u.image,
		t.specialty,
		t.description
	FROM patients p
	JOIN therapists t ON t.id = p.therapist_id
	JOIN users u ON u.id = t.user_id
	WHERE p.user_id = $1;
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

func (r *UserRepository) GetPatientPsychiatrist(userID int) (models.DoctorWithUser, error) {
	var psychiatrist models.DoctorWithUser

	query := `
	SELECT 
		p.id,
		u.name,
		u.email,
		u.age,
		u.image,
		p.description
	FROM patients pat
	JOIN psychiatrists p ON p.id = pat.psychiatrist_id
	JOIN users u ON u.id = p.user_id
	WHERE pat.user_id = $1;
	`

	row := r.DB.QueryRow(query, userID)
	err := row.Scan(
		&psychiatrist.ID,
		&psychiatrist.Name,
		&psychiatrist.Email,
		&psychiatrist.Age,
		&psychiatrist.Image,
		&psychiatrist.Description,
	)

	if err != nil {
		return models.DoctorWithUser{}, err
	}

	return psychiatrist, nil
}

func (r *UserRepository) GetAgendaByID(agendaID int) (models.Agenda, error){
	var agenda models.Agenda

	query := `
		SELECT id, professional_id, reserved
		FROM agendas
		WHERE id = $1
	`

	row := r.DB.QueryRow(query, agendaID)
	err := row.Scan(
		&agenda.ID,
		&agenda.ProfessionalID,
		&agenda.Reserved,
	)
	if err != nil {
		return models.Agenda{}, err
	}  
	return agenda, nil
}

func (r *UserRepository) GetTherapistPrice(therapistID int) (float64, error) {
	var price float64

	err := r.DB.QueryRow(`
		SELECT price
		FROM therapists
		WHERE id = $1
	`, therapistID).Scan(&price)

	return price, err
}

func (r *UserRepository) GetPsychiatristPrice(psychiastridID int) (float64, error) {
	var price float64
	err := r.DB.QueryRow(
		`SELECT price
			FROM psychiatrists WHERE id = $1
			`, psychiastridID).Scan(&price)
	if err != nil {
		return 0, err
	}

	return price, nil
}
