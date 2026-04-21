package models

type UserResponse struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Age int `json:"age"`
	Role string `json:"role"`
	Image string `json:"image"`
	Token string `json:"token"`
}

type TherapistResponse struct {
	Specialty string `json:"specialty"`
	Description string `json:"description"`
}

type PsychiatristResponse struct {
	Crm string `json:"crm"`
	Description string `json:"description"`
}

type PatientWithTherapistResponse struct {
	Name      string              `json:"name"`
	Therapist *TherapistInfo `json:"therapist"`
}

type PatientWithPsychiatristResponse struct {
	Name      string              `json:"name"`
	Psychiatrist *PsychiatristInfo `json:"psychiatrist"`
}

type DoctorDetailResponse struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Age int `json:"age"`
	Image string `json:"image"`
 
	Specialty string `json:"specialty"`
	CRM string `json:"crm"`
	Description string `json:"description"`
	Agendas []Agenda `json:"agendas"`
}

type UpdatePriceRequest struct {
	Price float64 `json:"price"`
}