package models

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Cpf string `json:"cpf"`
	Password string `json:"*"`
	Age int `json:"age"`
	Role string `json:"role"`
}

type PatientProfile struct{
	ID int `json:"id"`
	UserID int `json:"user_id"`
}

type TherapistProfile struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	Specialty string `json:"specialty"`
}

type PsychiatristProfile struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	CRM       string `json:"crm"`
}

