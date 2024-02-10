package http

import (
	"encoding/json"
	"net/http"

	"github.com/e2dk4r/supermarket"
)

// jsonPaginatedResponse writes data as json with page, perPage and total properties
func jsonPaginatedResponse(w http.ResponseWriter, page int, perPage int, total int, data interface{}) {
	jsonResponse(w, http.StatusOK, struct {
		Page    int         `json:"page"`
		PerPage int         `json:"perPage"`
		Total   int         `json:"total"`
		Data    interface{} `json:"data,omitempty"`
	}{
		Page:    page,
		PerPage: perPage,
		Total:   total,
		Data:    data,
	})
}

// jsonSuccessResponseMessage writes message with http status 200
func jsonSuccessResponseMessage(w http.ResponseWriter, message string) {
	jsonResponse(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{
		Message: message,
	})
}

// jsonSuccessResponse writes data with http status 200
func jsonSuccessResponse(w http.ResponseWriter, data interface{}) {
	jsonResponse(w, http.StatusOK, data)
}

// jsonFailResponse writes error with http status
func jsonFailResponse(w http.ResponseWriter, status int, err error) {
	jsonResponse(w, status, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}

// jsonResponse writes model as json to output
func jsonResponse(w http.ResponseWriter, status int, model interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(model); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(supermarket.ErrInternalServerError.Error()))
	}
}
