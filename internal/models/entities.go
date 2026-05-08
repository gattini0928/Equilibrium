package models

import "time"

type User struct {
	ID int
	Name string
	Email string
	Password string
	Age int
	Cpf string 
	Role string
	Image string
}

type UserPerfil struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Age int `json:"age"`
	Role string `json:"role"`
	Image string `json:"image"`
	CurrentDiagnosis string `json:"current_diagnosis"`
	Specialty   string `json:"specialty"`
	CRM string `json:"crm"`
	Description string `json:"description"`
	Price float64 `json:"price"`
}

type Patient struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	TherapistID      *int   `json:"therapist_id"`
	PsychiatristID   *int   `json:"psychiatrist_id"`
	CurrentDiagnosis string `json:"current_diagnosis"`
	Therapist *TherapistInfo `json:"therapist_info"`
	Psychiatrist *PsychiatristInfo `json:"psychiatrist_info"`
}

type Therapist struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Specialty   string `json:"specialty"`
	Description string `json:"description"`
	Price float64 `json:"price"`
}

type Psychiatrist struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	CRM    string `json:"crm"`
	Description string `json:"description"`
	Price float64 `json:"price"`
}

type TherapistInfo struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Age int `json:"age"`
	Image string `json:"image"`
	Specialty string `json:"specialty"`
	Description string `json:"description"`
}

type PsychiatristInfo struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Age int `json:"age"`
	Image string `json:"image"`
	Description string `json:"description"`
}

type Consultation struct {
	ID               int     `json:"id"`
	PatientID        int     `json:"patient_id"`
	TherapistID   int     `json:"therapist_id"`
	PsychiatristID string  `json:"psychiatrist_id"`
	Date             time.Time  `json:"date"`
	Price            float64 `json:"price"`
	Annotation       string  `json:"annotation"`
	AgendaID int `json:"agenda_id"`
	Status 			string `json:"status"`
}

type ConsultationDetail struct {
	ID              int
	PatientID       int
	TherapistID     *int
	PsychiatristID  *int
	Date            time.Time
	Price           float64
	Annotation      string
	Status          string

	Books    []Book
	Remedies []Remedy
}

type Book struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
}

type Remedy struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Dosage   string `json:"dosage"`
	Quantity int    `json:"quantity"`
}

type Agenda struct {
	ID             int    `json:"id"`
	ProfessionalID int    `json:"professional_id"`
	Day            int    `json:"day"`
	Month          int    `json:"month"`
	Hour           string `json:"hour"`
	Reserved       bool   `json:"reserved"`
}

type DoctorWithUser struct {
	ID          int
	Name        string
	Email       string
	Image       string
	Age         int
	Specialty   string
	CRM 		string
	Description string
	Price float64
}

type PatientWithUser struct {
	ID          int
	Name        string
	Email       string
	Image       string
	Age         int
	CurrentDiagnosis string
	Books       []Book
	Remedies []Remedy
}


type DoctorDashboard struct {
	Perfil   UserPerfil   `json:"perfil"`
	Agendas  []Agenda     `json:"agendas"`
	Patients []PatientWithUser
	Consultations []ConsultationDetail
}

