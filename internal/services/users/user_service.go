package users

import (
	"errors"

	"github.com/gattini0928/Equilibrium/internal/configs"
	"github.com/gattini0928/Equilibrium/internal/models"
	"github.com/gattini0928/Equilibrium/internal/services/auth"
	"github.com/gattini0928/Equilibrium/internal/services/validators"
)

var (
	ErrInvalidInput = errors.New("invalid input")
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidPassword = errors.New("invalid password")
	ErrTokenFailed = errors.New("token failed")
)

func (s *UserService) validateConsultationAccess(userID int, c models.Consultation, role string) error {
	switch role {

	case "therapist":
		id, err := s.Repo.GetTherapistIDByUserID(userID)
		if err != nil {
			return err
		}
		if c.TherapistID == nil || *c.TherapistID != id {
			return errors.New("forbidden")
		}

	case "psychiatrist":
		id, err := s.Repo.GetPsychiatristIDByUserID(userID)
		if err != nil {
			return err
		}
		if c.PsychiatristID == nil || *c.PsychiatristID != id {
			return errors.New("forbidden")
		}

	case "patient":
		id, err := s.Repo.GetPatientIDByUserID(userID)
		if err != nil {
			return err
		}
		if c.PatientID != id {
			return errors.New("forbidden")
		}

	default:
		return errors.New("invalid role")
	}

	return nil
}

func (s *UserService) CreateUser(user models.User, patient models.Patient, therapist models.Therapist, psychiatrist models.Psychiatrist) (string, error) {
	var err error

	err = validators.ValidateName(user.Name)
	if err != nil {
		return "", err
	}
	err = validators.ValidateEmail(user.Email)
	if err != nil {
		return "", err
	}

	err = validators.ValidatePassword(user.Password)
	if err != nil {
		return "",err
	}

	if user.Role == "therapist" || user.Role == "psychiatrist" {
		err = validators.ValidateAge(user.Age, user.Role)
		if err != nil {
			return "", err
		}
	}

	err = validators.ValidateCpf(user.Cpf)
	if err != nil {
		return "", err
	}

	hashPassword, err := validators.HashPassword(user.Password)
	if err != nil {
		return "", ErrInvalidPassword
	}

	user.Password = hashPassword

	err = s.Repo.CreateUserWithProfile(&user, &patient, &therapist, &psychiatrist)
	if err != nil {
		return "", err
	}

	cfg := configs.LoadDBConfig()

	token, err := auth.CreateJWT(s.Secret, user.ID, cfg.JWTExpirationInSeconds)
	if err != nil {
		return "", ErrTokenFailed
	}

	return token, nil

}

func (s *UserService) Login(email string, password string) (models.User, string, error) {
	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		return models.User{}, "", err
	}

	if !validators.CheckPasswordHash(password, user.Password){
		return models.User{}, "", ErrInvalidPassword
	}

	cfg := configs.LoadDBConfig()

	token, err := auth.CreateJWT(s.Secret, user.ID, cfg.JWTExpirationInSeconds)
	if err != nil {
		return models.User{}, "", ErrTokenFailed
	}

	return user, token, nil
}

// Completar cadastro do therapeuta
func (s *UserService) CompleteTherapistSignUp(userID int, specialty string, description string) error {
	err := validators.ValidateSpecialty(specialty)
	if err != nil {
		return err
	}

	err = validators.ValidateDescription(description)
	if err != nil {
		return err
	}

	return s.Repo.CompleteTherapist(userID, specialty, description)
}

// Completar cadastro do psiquiatra
func (s *UserService) CompletePsychiatristSignUp(userID int, crm string, description string) error {
	err := validators.ValidateCrm(crm)
	if err != nil {
		return err
	}

	err = validators.ValidateDescription(description)
	if err != nil {
		return err
	}

	return s.Repo.CompletePsychiatrist(userID, crm, description)
}

// Listagem de todos terapeutas
func (s *UserService) ListAllTherapists() ([]models.DoctorWithUser, error) {
	return s.Repo.GetAllTherapists()
}

// Listagem de todos psiquiatras
func (s *UserService) ListAllPsychiatrists() ([]models.DoctorWithUser, error) {
	return s.Repo.GetAllPsychiatrists()
}

// Detalhes do terapeuta
func (s *UserService) TherapistDetail(userID int) (models.DoctorWithUser, []models.Agenda, error) {

	therapist, err := s.Repo.GetTherapistById(userID)
	if err != nil {
		return models.DoctorWithUser{}, nil, err
	}

	agendas, err := s.Repo.GetTherapistAgenda(therapist.ID)
	if err != nil {
		return models.DoctorWithUser{}, nil, err
	}

	return therapist, agendas, nil
}

// Detalhes do psiquiatra
func (s *UserService) PsychiatristDetail(userID int) (models.DoctorWithUser, []models.Agenda, error) {

	psychiatrist, err := s.Repo.GetPsychiatristById(userID)
	if err != nil {
		return models.DoctorWithUser{}, nil, err
	}

	agendas, err := s.Repo.GetTherapistAgenda(psychiatrist.ID)
	if err != nil {
		return models.DoctorWithUser{}, nil, err
	}

	return psychiatrist, agendas, nil
}

func (s *UserService) TherapistToPatient(patientID, therapistID int) error {
	user, err := s.Repo.GetUserByID(patientID)

	if err != nil {
		return err
	}

	if user.Role != "patient" {
		return errors.New("forbidden")
	}
	
	return s.Repo.AddTherapistToPatient(patientID, therapistID)
}

func (s *UserService) PsychiatristToPatient(patientID, psychiatristID int) error {
	user, err := s.Repo.GetUserByID(patientID)
	if err != nil {
		return err
	}

	if user.Role != "patient" {
		return errors.New("forbidden")
	}

	return s.Repo.AddPsychiatristToPatient(patientID, psychiatristID)

}

// Terapeuta ou Psiquiatra vê os detalhes do paciente
func (s *UserService) GetPatientDetail(patientID, doctorID int) (models.PatientWithUser, error) {
	user, err := s.Repo.GetUserByID(doctorID)
	if err != nil {
		return models.PatientWithUser{}, err
	}

	if user.Role == "therapist" {
		return s.Repo.GetTherapistPatient(patientID)
	}

	if user.Role == "psychiatrist" {
		return s.Repo.GetPsychiatristPatient(patientID)
	}

	return models.PatientWithUser{}, errors.New("forbidden")
}	

// Detalhes do Terapeuta do Paciente
func (s *UserService) PatientTherapistDetail(userID int) (models.DoctorWithUser, error) {
	return s.Repo.GetPatientTherapist(userID)
}

// Detalhes do Psiquiatra do Paciente
func (s *UserService) PatientPsiquiatristDetail(userID int) (models.DoctorWithUser, error) {
	return s.Repo.GetPatientPsychiatrist(userID) 
}

// Listar pacientes do terapeuta ou psiquiatra
func (s *UserService) ListMyPatients(userID int) ([]models.PatientWithUser, error) {
	user, err := s.Repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	switch user.Role {
	case "therapist":
		return s.Repo.GetTherapistPatients(userID)

	case "psychiatrist":
		return s.Repo.GetPsychiatristPatients(userID)

	default:
		return nil, errors.New("forbidden")
	}
}

// Funções para agenda e preço
func (s *UserService) AddAgenda(userID int, day int, month int, hour string) (models.Agenda, error) {
	user, err := s.Repo.GetUserByID(userID)
	if err != nil {
		return models.Agenda{}, err
	}

	var professionalID int

	switch user.Role {
	case "therapist":
		professionalID, err = s.Repo.GetTherapistIDByUserID(userID)
		if err != nil {
			return models.Agenda{}, err
		}
	case "psychiatrist":
		professionalID, err = s.Repo.GetPsychiatristIDByUserID(userID)
		if err != nil {
			return models.Agenda{}, err
		}
	default:
		return models.Agenda{}, errors.New("forbidden")
	}

	return s.Repo.InsertAgenda(professionalID, day, month, hour)
}

func (s *UserService) RemoveAgenda(userID int, agendaID int) error {
	return s.Repo.DeleteAgenda(userID, agendaID)
}

func (s *UserService) UpdatePrice(userID int, price float64) error {
	user, err := s.Repo.GetUserByID(userID)
	if err != nil {
		return err
	}

	switch user.Role {
	case "therapist":
		return s.Repo.UpdateTherapistPrice(userID, price)

	case "psychiatrist":
		return s.Repo.UpdatePsychiatristPrice(userID, price)

	default:
		return errors.New("invalid role")
	}
}

func (s *UserService) Perfil(userID int) (any, error) {
	user, err := s.Repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	switch user.Role {
	case "patient":
		return s.Repo.GetPatientPerfil(userID)

	case "therapist":
		perfil, err := s.Repo.GetTherapistPerfil(userID)
		if err != nil {
			return nil, err
		}
		agendas, err := s.Repo.GetTherapistPrivateAgenda(userID)
		if err != nil {
			return nil, err
		}

		patientID, err := s.Repo.GetPatientIDByUserID(userID)
		if err != nil {
			return nil, err
		}

		patients,  err := s.Repo.GetTherapistPatients(patientID)
		if err != nil {
			return nil, err
		}

		therapistID, err := s.Repo.GetTherapistIDByUserID(userID)
		if err != nil {
			return nil, err
		}

		consultations, err := s.Repo.GetTherapistConsultations(therapistID)
		if err != nil {
			return nil, err
		}

		return models.DoctorDashboard{
			Perfil: perfil,
			Agendas: agendas,
			Patients: patients,
			Consultations: consultations,
		}, nil

	case "psychiatrist":
		perfil, err := s.Repo.GetPsychiatristPerfil(userID)
		if err != nil {
			return nil, err
		}
		agendas, err := s.Repo.GetPsychiatristPrivateAgenda(userID)
		if err != nil {
			return nil, err
		}

		patientID, err := s.Repo.GetPatientIDByUserID(userID)
		if err != nil {
			return nil, err
		}

		patients,  err := s.Repo.GetPsychiatristPatients(patientID)
		if err != nil {
			return nil, err
		}

		therapistID, err := s.Repo.GetPsychiatristIDByUserID(userID)
		if err != nil {
			return nil, err
		}

		consultations, err := s.Repo.GetPsychiatristConsultations(therapistID)
		if err != nil {
			return nil, err
		}
		return models.DoctorDashboard{
			Perfil: perfil,
			Agendas: agendas,
			Patients: patients,
			Consultations: consultations,
		}, nil

	default:
		return nil, errors.New("invalid role")
	}
}

func (s *UserService) ReserveTherapistAgenda(patientUserID, therapistID, agendaID int) error {
	patientID, err := s.Repo.GetPatientIDByUserID(patientUserID)
	if err != nil {
		return err
	}

	agenda, err := s.Repo.GetAgendaByID(agendaID)
	if err != nil {
		return err
	}

	if agenda.Reserved {
		return errors.New("agenda já reservada")
	}

	if agenda.ProfessionalID != therapistID {
		return errors.New("agenda inválida")
	}

	price, err := s.Repo.GetTherapistPrice(therapistID)
	if err != nil {
		return err
	}

	tx, err := s.Repo.DB.Begin()
	if err != nil {
		return err
	}

	err = s.Repo.MarkAgendaReserved(tx, agendaID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = s.Repo.CreateTherapistConsultation(tx, patientID, therapistID, agendaID, price)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (s *UserService) ReservePsychiatristAgenda(patientUserID, psychiatristID, agendaID int) error {
	patientID, err := s.Repo.GetPatientIDByUserID(patientUserID)
	if err != nil {
		return err
	}

	agenda, err := s.Repo.GetAgendaByID(agendaID)
	if err != nil {
		return err
	}

	if agenda.Reserved {
		return errors.New("agenda já reservada")
	}

	if agenda.ProfessionalID != psychiatristID {
		return errors.New("agenda inválida")
	}

	price, err := s.Repo.GetPsychiatristPrice(psychiatristID)
	if err != nil {
		return err
	}

	tx, err := s.Repo.DB.Begin()
	if err != nil {
		return err
	}

	err = s.Repo.CreatePsychiatristConsultation(tx, patientID, psychiatristID, agendaID, price)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = s.Repo.MarkAgendaReserved(tx, agendaID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (s *UserService) ShowConsultations(userID int) ([]models.Consultation, error) {
	user, err := s.Repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	switch user.Role {
	case "patient":
		patientID, err := s.Repo.GetPatientIDByUserID(userID)
		if err != nil {
			return nil, err
		}
		return s.Repo.GetPatientConsultations(patientID)
	case "therapist":
		therapistID, err := s.Repo.GetTherapistIDByUserID(userID)
		if err != nil {
			return nil, err
		}
		return s.Repo.GetTherapistConsultations(therapistID)
	case "psychiatrist":
		psychiatristID, err := s.Repo.GetPsychiatristIDByUserID(userID)
		if err != nil {
			return nil, err
		}
		return s.Repo.GetPsychiatristConsultations(psychiatristID)
	}
	return nil, errors.New("invalid role")
}

func (s *UserService) ShowConsultation(userID, consultationID int) (models.Consultation, error) {
	user, err := s.Repo.GetUserByID(userID)
	if err != nil {
		return models.Consultation{}, err
	}

	c, err := s.Repo.GetConsultationByID(consultationID)
	if err != nil {
		return models.Consultation{}, err
	}

	switch user.Role {

	case "patient":
		patientID, err := s.Repo.GetPatientIDByUserID(userID)
		if err != nil {
			return models.Consultation{}, err
		}
		if c.PatientID != patientID {
			return models.Consultation{}, errors.New("forbidden")
		}

	case "therapist":
		therapistID, err := s.Repo.GetTherapistIDByUserID(userID)
		if err != nil {
			return models.Consultation{}, err
		}
		if c.TherapistID == nil || *c.TherapistID != therapistID {
			return models.Consultation{}, errors.New("forbidden")
		}

	case "psychiatrist":
		psychiatristID, err := s.Repo.GetPsychiatristIDByUserID(userID)
		if err != nil {
			return models.Consultation{}, err
		}
		if c.PsychiatristID == nil || *c.PsychiatristID != psychiatristID {
			return models.Consultation{}, errors.New("forbidden")
		}

	default:
		return models.Consultation{}, errors.New("invalid role")
	}

	return c, nil
}

func (s *UserService) StartConsultation(userID, consultationID int) error {
	c, err := s.Repo.GetConsultationByID(consultationID)
	if err != nil {
		return err
	}

	user, err := s.Repo.GetUserByID(userID)
	if err != nil {
		return err
	}

	err = s.validateConsultationAccess(userID, c, user.Role)
	if err != nil {
		return err
	}

	return s.Repo.UpdateConsultationInProgress(consultationID)
}

func (s *UserService) FinishConsultation(userID, consultationID int) error {
	c, err := s.Repo.GetConsultationByID(consultationID)
	if err != nil {
		return err
	}

	user, err := s.Repo.GetUserByID(userID)
	if err != nil {
		return err
	}

	err = s.validateConsultationAccess(userID, c, user.Role)
	if err != nil {
		return err
	}

	return s.Repo.UpdateConsultationFinished(consultationID)
}

func (s *UserService) SaveConsultationRemedy(userID, consultationID int, remedyName, remedyDosage string, remedyQuantity int) error {
	c, err := s.Repo.GetConsultationByID(consultationID)
	if err != nil {
		return err
	}

	user, err := s.Repo.GetUserByID(userID)
	if err != nil {
		return err
	}

	switch user.Role {
	case "psychiatrist":
		id, err := s.Repo.GetPsychiatristIDByUserID(userID)
		if err != nil {
			return err
		}

		if c.PsychiatristID == nil || *c.PsychiatristID != id {
			return errors.New("forbidden")
		}

		remedyID, err := s.Repo.InsertRemedy(remedyName, remedyDosage, remedyQuantity)
		if err != nil {
			return err
		}

		return s.Repo.LinkRemedyToConsultation(consultationID, remedyID)

	default:
		return errors.New("forbidden")
	}
}

func (s *UserService) SaveConsultationBook(userID, consultationID int, author, title string) error {
	c, err := s.Repo.GetConsultationByID(consultationID)
	if err != nil {
		return err
	}

	user, err := s.Repo.GetUserByID(userID)
	if err != nil {
		return err
	}

	switch user.Role {
	case "therapist":
		id, err := s.Repo.GetTherapistIDByUserID(userID)
		if err != nil {
			return err
		}

		if c.TherapistID == nil || *c.TherapistID != id {
			return errors.New("forbidden")
		}

		bookID, err := s.Repo.InsertBook(author, title)
		if err != nil {
			return err
		}

			return s.Repo.LinkBookToConsultation(consultationID, bookID)
	default:
		return errors.New("forbidden")
	}
}

func (s *UserService) SaveConsultationAnnotation(userID, consultationID int, annotation string) error {
	c, err := s.Repo.GetConsultationByID(consultationID)
	if err != nil {
		return err
	}

	user, err := s.Repo.GetUserByID(userID)
	if err != nil {
		return err
	}

	err = s.validateConsultationAccess(userID, c, user.Role)
	if err != nil {
		return err
	}

	return s.Repo.UpdateAnnotationConsultation(consultationID, annotation)
}
