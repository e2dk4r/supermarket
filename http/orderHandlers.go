package http

import (
	"fmt"
	"net/http"
	"path/filepath"
)

func (h *Handler) OrderIndex(w http.ResponseWriter, r *http.Request) {
	orderId := filepath.Base(r.URL.Path)

	jsonSuccessResponse(w, struct {
		Id string `json:"id,omitempty"`
	}{
		Id: orderId,
	})
}

func (h *Handler) OrderShow(w http.ResponseWriter, r *http.Request) {
	jsonFailResponse(w, http.StatusInternalServerError, fmt.Errorf("not implemented"))
}

func (h *Handler) OrderCreate(w http.ResponseWriter, r *http.Request) {
	jsonFailResponse(w, http.StatusInternalServerError, fmt.Errorf("not implemented"))
}
