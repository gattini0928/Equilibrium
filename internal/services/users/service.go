package users

import (
    repoUsers "github.com/gattini0928/Equilibrium/internal/repositories/users"

)

type UserService struct {
    Repo *repoUsers.UserRepository
	Secret []byte
}


func NewUserService(r *repoUsers.UserRepository, secret []byte) *UserService {
    return &UserService{Repo: r, Secret: secret}
}
