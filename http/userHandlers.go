package http

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/e2dk4r/supermarket"
)

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	random, err := h.RandomService.GenerateString(32)
	if err != nil {
		jsonFailResponse(w, http.StatusInternalServerError, supermarket.ErrInternalServerError)
		return
	}

	jsonSuccessResponse(w, struct {
		Message string
	}{
		Message: random,
	})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	defer r.Body.Close()

	{
		reader := io.LimitReader(r.Body, 1<<20)
		err := json.NewDecoder(reader).Decode(&user)
		if err != nil {
			log.Print(err)
			jsonFailResponse(w, http.StatusBadRequest, errors.New("cannot parse request"))
			return
		}
	}
	if user.Username == "" || len(user.Username) < 3 || len(user.Password) > 256 {
		jsonFailResponse(w, http.StatusUnprocessableEntity, errors.New("username is required. min 3 max 256"))
		return
	}
	if user.Password == "" || len(user.Password) < 8 || len(user.Password) > 256 {
		jsonFailResponse(w, http.StatusUnprocessableEntity, errors.New("password is required. min 8 max 256"))
		return
	}

	token, err := h.AuthService.CreateToken(&supermarket.User{Username: user.Username, Password: user.Password})
	if err != nil {
		if err == supermarket.ErrPasswordNotMatch {
			jsonFailResponse(w, http.StatusUnauthorized, errors.New("username or password is not correct"))
			return
		}

		jsonFailResponse(w, http.StatusInternalServerError, supermarket.ErrInternalServerError)
		return
	}

	jsonSuccessResponse(w, token)
}
