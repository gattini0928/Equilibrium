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
	General string
}

type PsychiatristForm struct {
	CRM string
	Description string
	Errors map[string]string
	General string
}

type ViewData struct {
    IsAuth bool
}

type LoginView struct {
    ViewData
    Form LoginForm
	Msg string
}

type SignupView struct {
    ViewData
    Form SignupForm
}

type TherapistView struct {
	ViewData
    Form TherapistForm
	Msg string
}

type PsychiatristView struct {
	ViewData
    Form PsychiatristForm
	Msg string
}

type PerfilView struct {
	ViewData
	Messages map[string]string
	Perfil any	
}

type StatusView struct {
	ViewData
	StatusCode int
	StatusMessage string

}