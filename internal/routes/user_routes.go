package routes

import (
	"net/http"

	handlerUsers "github.com/gattini0928/Equilibrium/internal/handlers/users"
	"github.com/gattini0928/Equilibrium/internal/services/auth"
)

func UserRoutes(mux *http.ServeMux, h *handlerUsers.UserHandler, secret []byte) {
	mux.HandleFunc("POST /signup", h.HandleSignup)
	mux.HandleFunc("POST /login", h.HandleLogin)
	mux.Handle("POST /complete-therapist-detail",
		auth.JWTMiddleware(secret, http.HandlerFunc(h.HandleCompleteTherapist)))
	mux.Handle("POST /complete-psychiatrist-detail",
		auth.JWTMiddleware(secret, http.HandlerFunc(h.HandleCompletePsychiatrist)))
	mux.HandleFunc("GET /therapists", h.HandleAllTherapists)
	mux.HandleFunc("GET /psychiatrists", h.HandleAllPsychiatrists)
}
