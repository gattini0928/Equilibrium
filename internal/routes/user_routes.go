package routes

import (
	"net/http"

	handlerUsers "github.com/gattini0928/Equilibrium/internal/handlers/users"
	"github.com/gattini0928/Equilibrium/internal/services/auth"
)


func UserRoutes(mux *http.ServeMux, h *handlerUsers.UserHandler, secret []byte) {

	mux.Handle("/static/",
	http.StripPrefix("/static/",
		http.FileServer(http.Dir("static")),
		),
	)			

	mux.HandleFunc("GET /{$}", h.HandleHome)

	// AUTH
	mux.HandleFunc("POST /signup", h.HandleSignup)
	mux.HandleFunc("POST /login", h.HandleLogin)

	// COMPLETAR PERFIL (JWT)
	mux.Handle("POST /therapists/profile",
		auth.JWTMiddleware(secret, http.HandlerFunc(h.HandleCompleteTherapist)))
	mux.Handle("POST /psychiatrists/profile",
		auth.JWTMiddleware(secret, http.HandlerFunc(h.HandleCompletePsychiatrist)))

	// Perfil
	mux.Handle("GET /me",
	auth.JWTMiddleware(secret, http.HandlerFunc(h.HandlePerfil)))

	// LISTAGEM PÚBLICA
	mux.HandleFunc("GET /therapists", h.HandleAllTherapists)
	mux.HandleFunc("GET /psychiatrists", h.HandleAllPsychiatrists)

	// DETALHES(Clique no card)
	mux.HandleFunc("GET /therapists/{id}", h.HandleTherapistDetail)
	mux.HandleFunc("GET /psychiatrists/{id}", h.HandlePsychiatristDetail)

	mux.Handle("POST /therapists/{therapist_id}/agendas/{agenda_id}/reserve", 
		auth.JWTMiddleware(secret, http.HandlerFunc(h.HandleReserveTherapistAgenda)))
	mux.Handle("POST /psychiatrists/{psychiatrist_id}/agendas/{agenda_id}/reserve", 
		auth.JWTMiddleware(secret, http.HandlerFunc(h.HandleReservePsychiatristAgenda)))

	// VINCULAR (JWT) PACIENTE -> SEU TERAPEUTA - PSIQUIATRA
	mux.Handle("PUT /me/therapist/{id}",
		auth.JWTMiddleware(secret, http.HandlerFunc(h.HandleAddTherapistToPatient)))
	mux.Handle("PUT /me/psychiatrist/{id}",
		auth.JWTMiddleware(secret, http.HandlerFunc(h.HandleAddPsychiatristToPatient)))

	// PEFIL LOGADO (JWT) PACIENTE -> SEU TERAPEUTA
	mux.Handle("GET /me/therapist",
		auth.JWTMiddleware(secret, http.HandlerFunc(h.HandlePatientTherapistDetail)))
	mux.Handle("GET /me/psychiatrist",
		auth.JWTMiddleware(secret, http.HandlerFunc(h.HandlePatientPsychiatristDetail)))

	// TODOS PACIENTES TERAPEUTA / PSIQUIATRA
	mux.Handle("GET /me/patients",
		auth.JWTMiddleware(secret, http.HandlerFunc(h.HandleMyPatients)))

	// DETALHE DO PACIENTE (JWT)
	mux.Handle("GET /patients/{id}",
		auth.JWTMiddleware(secret, http.HandlerFunc(h.HandlePatientDetail)))
	
	// Manipulação de Agenda
	mux.Handle("POST /me/agenda",
		auth.JWTMiddleware(secret, http.HandlerFunc(h.HandleAddAgenda)))
	mux.Handle("DELETE /me/agenda/{agenda_id}",
		auth.JWTMiddleware(secret, http.HandlerFunc(h.HandleDeleteAgenda)))
	mux.Handle("PUT /me/price",
		auth.JWTMiddleware(secret, http.HandlerFunc(h.HandleUpdatePrice)))

	mux.Handle("GET /logout",
		auth.JWTMiddleware(secret, http.HandlerFunc(h.HandleLogout)))
}
