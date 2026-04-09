package users

import (
	userService "github.com/gattini0928/Equilibrium/internal/services/users"
)

type UserHandler struct {
	Service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{Service: service}
}