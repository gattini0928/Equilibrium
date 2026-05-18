package routes


import (
	"net/http"

	handlerUsers "github.com/gattini0928/Equilibrium/internal/handlers/users"
	"github.com/gattini0928/Equilibrium/internal/services/auth"
)

func ConsultationRoutes(mux *http.ServeMux, h *handlerUsers.UserHandler, secret []byte) {
	// Entrar na consulta
	mux.Handle("GET /consultations",
		auth.JWTMiddleware(secret, http.HandlerFunc(h.HandleConsultation)))
	// Atualizar para progresso
	mux.Handle("POST /consultations/from-agenda/{agenda_id}/start",
		auth.JWTMiddleware(secret, http.HandlerFunc(h.HandleStartConsultation)))
	// Salvar Infos da Consulta 
	mux.Handle("PUT /consultations/{consultation_id}",
		auth.JWTMiddleware(secret, http.HandlerFunc(h.HandleSaveConsultationInfos)))
	// Detalhes da consulta
	mux.Handle("GET /consultations/{consultation_id}",
		auth.JWTMiddleware(secret, http.HandlerFunc(h.HandleConsultationDetail)))
}