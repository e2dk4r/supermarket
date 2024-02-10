package http

import (
	"fmt"
	"net/http"
)

func (h *Handler) ProductIndex(w http.ResponseWriter, r *http.Request) {
	jsonFailResponse(w, http.StatusInternalServerError, fmt.Errorf("not implemented"))
}

func (h *Handler) ProductShow(w http.ResponseWriter, r *http.Request) {
	jsonFailResponse(w, http.StatusInternalServerError, fmt.Errorf("not implemented"))
}

func (h *Handler) ProductCreate(w http.ResponseWriter, r *http.Request) {
	jsonFailResponse(w, http.StatusInternalServerError, fmt.Errorf("not implemented"))
}
