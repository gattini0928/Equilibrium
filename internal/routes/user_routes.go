package routes

import (
	"net/http"
	handlerUsers "github.com/gattini0928/Equilibrium/internal/handlers/users"
)

func UserRoutes(mux *http.ServeMux, h *handlerUsers.UserHandler) {
	mux.HandleFunc("POST /signup", h.HandleSignup)
	mux.HandleFunc("POST /login", h.HandleLogin)
	mux.HandleFunc("POST /complete-therapist-detail", h.HandleCompleteTherapist)
	mux.HandleFunc("POST /complete-psychiatrist-detail", h.HandleCompletePsychiatrist)
	mux.HandleFunc("GET /therapists", h.HandleAllTherapists)
	mux.HandleFunc("GET /psychiatrists", h.HandleAllPsychiatrists)
}
