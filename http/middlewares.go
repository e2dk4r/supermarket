package http

import (
	"errors"
	"log"
	"net/http"
	"strings"
)

func (h *Handler) MRequireAuth(w http.ResponseWriter, r *http.Request) bool {
	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	if token == "" {
		jsonFailResponse(w, http.StatusUnauthorized, errors.New("bearer token not found in authorization header"))
		return false
	}

	err := h.AuthService.VerifyToken(token)
	if err != nil {
		log.Print(err)
		jsonFailResponse(w, http.StatusUnauthorized, errors.New("token is not valid"))
		return false
	}

	return true
}
