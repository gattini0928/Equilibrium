package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"errors"
	"context"
	"github.com/gattini0928/Equilibrium/internal/services/auth"
)

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("Corpo da requisição vazio")
	}
	
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error":err.Error()})
}

func CheckID(pathID string, r *http.Request) (int, error) {
	idStr := r.PathValue(pathID)
	if idStr == "" {
		return 0, errors.New("id é obrigatório")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func CheckJWT(w http.ResponseWriter, ctx context.Context) (int, bool) {
	val := ctx.Value(auth.UserIDKey)

	userID, ok := val.(int)
	if !ok {
		WriteJSON(w, http.StatusUnauthorized, "não autorizado")
		return 0, false
	}

	return userID, true
}