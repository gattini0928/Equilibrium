package routes

import (
	"net/http"
	handlerUsers "github.com/gattini0928/Equilibrium/internal/handlers/users"
)

func UserRoutes(mux *http.ServeMux, h *handlerUsers.UserHandler) {
	mux.HandleFunc("POST /signup", h.HandleSignup)
	mux.HandleFunc("POST /login", h.HandleLogin)
}
