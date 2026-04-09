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
}

type Psychiatrist struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	CRM    string `json:"crm"`
	Description string `json:"description"`
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

type PatientWithTherapistResponse struct {
	Name      string              `json:"name"`
	Therapist *TherapistInfo `json:"therapist"`
}

type PatientWithPsychiatristResponse struct {
	Name      string              `json:"name"`
	Psychiatrist *PsychiatristInfo `json:"psychiatrist"`
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
	Date             string  `json:"date"`
	Price            float64 `json:"price"`
	Annotation       string  `json:"annotation"`
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

type DoctorWithUser struct {
	ID          int
	Name        string
	Email       string
	Image       string
	Age         int
	Specialty   string
	CRM 		string
	Description string
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

type DoctorDetailResponse struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Age int `json:"age"`
	Image string `json:"image"`
 
	Specialty string `json:"specialty"`
	CRM string `json:"crm"`
	Description string `json:"description"`
}

