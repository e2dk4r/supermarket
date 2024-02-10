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

func (h *Handler) MSecurityHeaders(w http.ResponseWriter, r *http.Request) bool {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Referrer-Policy", "no-referrer-when-downgrade")
	w.Header().Set("Content-Security-Policy", "default-src 'self' http: https: ws: wss: data: blob: 'unsafe-inline'; frame-ancestors 'self';")
	w.Header().Set("Permissions-Policy", "interest-cohort=()")
	w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
	return true
}
