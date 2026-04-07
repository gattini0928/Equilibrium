package users

import (
	"net/http"
	"errors"
	"github.com/gattini0928/Equilibrium/internal/models"
	serviceUsers "github.com/gattini0928/Equilibrium/internal/services/users"
	"github.com/gattini0928/Equilibrium/internal/utils"
	"github.com/gattini0928/Equilibrium/internal/services/auth"
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

	userID := r.Context().Value(auth.UserIDKey).(int)

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

	userID := r.Context().Value(auth.UserIDKey).(int)

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
		Role: user.Role,
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

