package users

import (
	userService "github.com/gattini0928/Equilibrium/internal/services/users"
)

type UserHandler struct {
	Service userService.UserServiceInterface
}

func NewUserHandler(s userService.UserServiceInterface ) *UserHandler {
	return &UserHandler{Service: s}
}