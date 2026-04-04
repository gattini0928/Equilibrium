package users

import (
	userService "github.com/gattini0928/Equilibrium/internal/services/users"
)

type UserHandler struct {
	Service *userService.UserService
}

func NewUserHandler(s *userService.UserService) *UserHandler{
	return &UserHandler{Service:s}
}