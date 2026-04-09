package users

import (
	"errors"
	"net/http"
	"strconv"
	"database/sql"

	"github.com/gattini0928/Equilibrium/internal/models"
	"github.com/gattini0928/Equilibrium/internal/services/auth"
	serviceUsers "github.com/gattini0928/Equilibrium/internal/services/users"
	"github.com/gattini0928/Equilibrium/internal/utils"
)

func (h *UserHandler) HandleSignup(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest

	err := utils.ParseJSON(r, &req)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user := models.User{
		Name: req.Name,
		Email: req.Email,
		Password: req.Password,
		Age: req.Age,
		Cpf: req.Cpf,
		Role: req.Role,
		Image: req.Image,
	}

	var patient models.Patient
	var therapist models.Therapist
	var psychiatrist models.Psychiatrist

	switch req.Role {
	case "therapist":
		therapist.Specialty = req.Specialty
		therapist.Description = req.Description

	case "psychiatrist":
		psychiatrist.CRM = req.CRM
		psychiatrist.Description = req.Description
	}

	err = h.Service.CreateUser(user, patient, therapist, psychiatrist)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{
		"message": "Usuário criado com sucesso",
	})
}


func (h *UserHandler) HandleCompleteTherapist(w http.ResponseWriter, r *http.Request) {
	var therapist models.Therapist
	err := utils.ParseJSON(r, &therapist)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	if therapist.Specialty == "" || therapist.Description == "" {
		utils.WriteJSON(w, http.StatusBadRequest, "especialidade e descrição são obrigatórios")
		return 
	}

	val := r.Context().Value(auth.UserIDKey)
	userID, ok := val.(int)
	if !ok {
		utils.WriteJSON(w, http.StatusUnauthorized, "não autorizado")
		return
	}

	err = h.Service.CompleteTherapistSignUp(userID, therapist.Specialty, therapist.Description)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	resp := models.TherapistResponse {
		Specialty: therapist.Specialty,
		Description: therapist.Description,
	}

	utils.WriteJSON(w, http.StatusOK, resp)
}

func (h *UserHandler) HandleCompletePsychiatrist(w http.ResponseWriter, r *http.Request) {
	var psychiatrist models.Psychiatrist
	err := utils.ParseJSON(r, &psychiatrist)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	if psychiatrist.CRM == "" || psychiatrist.Description == "" {
		utils.WriteJSON(w, http.StatusBadRequest, "CRM e descrição são obrigatórios")
		return 
	}

	val := r.Context().Value(auth.UserIDKey)
	userID, ok := val.(int)
	if !ok {
		utils.WriteJSON(w, http.StatusUnauthorized, "não autorizado")
		return
	}

	err = h.Service.CompletePsychiatristSignUp(userID, psychiatrist.CRM, psychiatrist.Description)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	resp := models.PsychiatristResponse {
		Crm: psychiatrist.CRM,
		Description: psychiatrist.Description,
	}

	utils.WriteJSON(w, http.StatusOK, resp)
}

func (h *UserHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var input models.LoginRequest
	err := utils.ParseJSON(r, &input)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user, token, err := h.Service.Login(input.Email, input.Password)
	if err != nil {

		if errors.Is(err, serviceUsers.ErrInvalidPassword) ||
			errors.Is(err, serviceUsers.ErrUserNotFound) {

			utils.WriteError(w, http.StatusUnauthorized, errors.New("email ou senha inválidos"))
			return
		}

		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := models.UserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Age: user.Age,
		Role: user.Role,
		Image: user.Image,
		Token: token,
	}

	utils.WriteJSON(w, http.StatusOK, res)
}

func (h *UserHandler) HandleAllTherapists(w http.ResponseWriter, r *http.Request) {

	therapists, err := h.Service.ListAllTherapists()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, therapists)
}

func (h *UserHandler) HandleAllPsychiatrists(w http.ResponseWriter, r *http.Request) {

	psychiatrists, err := h.Service.ListAllPsychiatrists()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, psychiatrists)
}

func (h *UserHandler) HandleTherapistDetail(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.PathValue("user_id")

	if userIdStr == "" {
		utils.WriteJSON(w, http.StatusBadGateway, "id é obrigatório")
		return
	}

	userID, err := strconv.Atoi(userIdStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	therapist, err := h.Service.TherapistDetail(userID)
	if errors.Is(err, sql.ErrNoRows) {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	resp := models.DoctorDetailResponse {
		Name: therapist.Name,
		Email: therapist.Email,
		Age: therapist.Age,
		Image: therapist.Image,
		Specialty: therapist.Specialty,
		Description: therapist.Description,
	}

	utils.WriteJSON(w, http.StatusOK, resp)
}


func (h *UserHandler) HandlePsychiatristDetail(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.PathValue("user_id")

	if userIdStr == "" {
		utils.WriteJSON(w, http.StatusBadRequest, "id é obrigatório")
		return
	}

	userID, err := strconv.Atoi(userIdStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	psychiatrist, err := h.Service.PsychiatristDetail(userID)
	if errors.Is(err, sql.ErrNoRows) {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	resp := models.DoctorDetailResponse {
		Name: psychiatrist.Name,
		Email: psychiatrist.Email,
		Age: psychiatrist.Age,
		Image: psychiatrist.Image,
		CRM: psychiatrist.CRM,
		Description: psychiatrist.Description,
	}

	utils.WriteJSON(w, http.StatusOK, resp)
}

func (h *UserHandler) HandleAddTherapistToPatient(w http.ResponseWriter, r *http.Request) {
	patientIdStr := r.PathValue("patient_id")
	therapistIdStr := r.PathValue("therapist_id")

	if patientIdStr == "" || therapistIdStr == "" {
		utils.WriteJSON(w, http.StatusBadRequest, "id é obrigatório")
		return
	}

	patientID, err := strconv.Atoi(patientIdStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	therapistID, err := strconv.Atoi(therapistIdStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = h.Service.TherapistToPatient(patientID, therapistID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{
	"message": "Terapeuta vinculado com sucesso",
	})
}

func (h *UserHandler) HandleAddPsychiatristToPatient(w http.ResponseWriter, r *http.Request) {
	patientIdStr := r.PathValue("patient_id")
	psychiatristIdStr := r.PathValue("psychiatrist_id")

	if patientIdStr == "" || psychiatristIdStr == "" {
		utils.WriteJSON(w, http.StatusBadRequest, "id é obrigatório")
		return
	}

	patientID, err := strconv.Atoi(patientIdStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	psychiatristID, err := strconv.Atoi(psychiatristIdStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = h.Service.PsychiatristToPatient(patientID, psychiatristID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{
	"message": "Therapist vinculado com sucesso",
	})
}

func (h *UserHandler) HandlePatientTherapistDetail(w http.ResponseWriter, r *http.Request) {
	id, err := utils.CheckID("user_id", r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	therapist, err := h.Service.PatientTherapistDetail(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, therapist)
}


func (h *UserHandler) HandlePatientPsychiatristDetail(w http.ResponseWriter, r *http.Request) {
	id, err := utils.CheckID("user_id", r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	psychiatrist, err := h.Service.PatientPsiquiatristDetail(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, psychiatrist)
}

func (h *UserHandler) HandleTherapistAllPatients(w http.ResponseWriter, r *http.Request){
	val := r.Context().Value(auth.UserIDKey)
	therapist_id, ok := val.(int)

	if !ok {
		utils.WriteJSON(w, http.StatusUnauthorized, "não autorizado")
		return
	}

	patients, err := h.Service.ListAllTherapistPatients(therapist_id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError,err)
		return 
	}

	utils.WriteJSON(w, http.StatusOK, patients)

}

func (h *UserHandler) HandlePsychiatristAllPatients(w http.ResponseWriter, r *http.Request){
	val := r.Context().Value(auth.UserIDKey)
	psychiatrist_id, ok := val.(int)

	if !ok {
		utils.WriteJSON(w, http.StatusUnauthorized, "não autorizado")
		return
	}

	patients, err := h.Service.ListAllPsychiatristPatients(psychiatrist_id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError,err)
		return 
	}

	utils.WriteJSON(w, http.StatusOK, patients)
}

func (h *UserHandler) HandleTherapistPatientDetail(w http.ResponseWriter, r *http.Request) {
	id, err := utils.CheckID("user_id", r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	patient, err := h.Service.TherapistPatientDetail(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, patient)

} 

func (h *UserHandler) HandlePsychiatristPatientDetail(w http.ResponseWriter, r *http.Request) {
	id, err := utils.CheckID("user_id", r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	patient, err := h.Service.PsychiatristPatientDetail(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, patient)

} 
