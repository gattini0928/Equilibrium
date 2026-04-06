package models

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
	CRM string `json:"crm"`
}

type UserResponse struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Role string `json:"role"`
	Image string `json:"image"`
	Token string `json:"token"`
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
}

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

