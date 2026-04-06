package users

import (
	"net/http"
	"errors"
	"github.com/gattini0928/Equilibrium/internal/models"
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

