package routes


import (
	"net/http"

	handlerUsers "github.com/gattini0928/Equilibrium/internal/handlers/users"
	"github.com/gattini0928/Equilibrium/internal/services/auth"
)

func ConsultationRoutes(mux *http.ServeMux, h *handlerUsers.UserHandler, secret []byte) {
	// Tela da consulta
	mux.Handle("GET /consultations",
		auth.JWTMiddleware(secret, http.HandlerFunc(h.HandleConsultation)))
	// Salvar Infos da Consulta 
	mux.Handle("PUT /consultations/{id}",
		auth.JWTMiddleware(secret, http.HandlerFunc(h.HandleSaveConsultationInfos)))
	// Finalizar Consulta
	mux.Handle("PUT /consultations/{id}/finish",
		auth.JWTMiddleware(secret, http.HandlerFunc(h.HandleFinishConsultation)))
	// Consultas no perfil
	mux.Handle("GET /me/consultations",
		auth.JWTMiddleware(secret, http.HandlerFunc(h.HandleAllConsultations)))
	// Detalhes da consulta
	mux.Handle("GET /consultations/{id}",
		auth.JWTMiddleware(secret, http.HandlerFunc(h.HandleConsultationDetail)))
}