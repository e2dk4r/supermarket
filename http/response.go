package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

// jsonSuccessRepsonse writes data with http status 200
func jsonSuccessResponse(w http.ResponseWriter, data interface{}) {
	jsonResponse(w, http.StatusOK, data, nil)
}

// jsonSuccessRepsonse writes error with http status
func jsonFailResponse(w http.ResponseWriter, status int, err error) {
	jsonResponse(w, status, nil, err)
}

// jsonResponse writes data or error as json
func jsonResponse(w http.ResponseWriter, status int, data interface{}, err error) {
	resp := Response{
		Data: data,
	}
	if err != nil {
		resp.Error = err.Error()
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	if err = json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("[ERR] encoding resp %v:%s", resp, err)
		fmt.Fprintf(w, "internal server error")
	}
}