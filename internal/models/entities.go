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
	ProfessionalID   int     `json:"professional_id"`
	ProfessionalType string  `json:"professional_type"`
	Date             time.Time  `json:"date"`
	Price            float64 `json:"price"`
	Annotation       string  `json:"annotation"`
	Status 			string `json:"status"`
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
	CurrentDiagnosis *string
	Books       []Book
	Remedies []Remedy
}

type CreateUserRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Age int `json:"age"`
	Cpf string `json:"cpf"`
	Role string `json:"role"`
	Image string `json:"image"`
 
	Specialty string `json:"specialty"`
	CRM string `json:"crm"`
	Description string `json:"description"`
}

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}


