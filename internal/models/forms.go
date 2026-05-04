package models 

type LoginForm struct {
	Email    string
	Password string
	Errors   map[string]string
	General  string
}

type SignupForm struct {
	Name string
	Email string
	Password string
	Age int
	Cpf string 
	Role string
	Image string
	Errors map[string]string
	General string
}

type TherapistForm struct {
	Specialty string
	Description string
	Errors map[string]string
}

type PsychiatristForm struct {
	CRM string
	Description string
	Errors map[string]string
}