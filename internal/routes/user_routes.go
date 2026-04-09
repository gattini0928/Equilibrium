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
	mux.HandleFunc("GET /therapist-detail/{user_id}", h.HandleTherapistDetail)
	mux.HandleFunc("GET /psychiatrist-detail/{user_id}", h.HandlePsychiatristDetail)
	mux.HandleFunc("PUT /add-therapist-patient/{patient_id}/{therapist_id}", h.HandleAddTherapistToPatient)
	mux.HandleFunc("PUT /add-psychiatrist-patient/{patient_id}/{psychiatrist_id}", h.HandleAddPsychiatristToPatient)
	mux.Handle("GET /therapist-patient-detail/{user_id}",
		auth.JWTMiddleware(secret, http.HandlerFunc(h.HandleTherapistPatientDetail)))
	mux.Handle("GET /psychiatrist-patient-detail/{user_id}",
		auth.JWTMiddleware(secret, http.HandlerFunc(h.HandlePsychiatristPatientDetail)))
	mux.Handle("GET /therapist-all-patients",
		auth.JWTMiddleware(secret, http.HandlerFunc(h.HandleTherapistAllPatients)))
	mux.Handle("GET /psychiatrist-all-patients",
		auth.JWTMiddleware(secret, http.HandlerFunc(h.HandlePsychiatristAllPatients)))
}
