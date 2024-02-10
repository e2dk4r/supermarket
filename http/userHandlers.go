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
	var request struct {
		Username             string `json:"username"`
		Password             string `json:"password"`
		PasswordConfirmation string `json:"password_confirmation"`
		Token                string `json:"token"`
	}

	// parse
	defer r.Body.Close()
	{
		reader := io.LimitReader(r.Body, 1<<20)
		err := json.NewDecoder(reader).Decode(&request)
		if err != nil {
			jsonFailResponse(w, http.StatusBadRequest, errors.New("cannot parse request"))
			return
		}
	}

	// validate
	if request.Username == "" || len(request.Username) < 3 || len(request.Username) > 256 {
		jsonFailResponse(w, http.StatusUnprocessableEntity, errors.New("username is required. min 3 max 256"))
		return
	}
	if request.Password == "" || len(request.Password) < 8 || len(request.Password) > 256 {
		jsonFailResponse(w, http.StatusUnprocessableEntity, errors.New("password is required. min 8 max 256"))
		return
	}
	if request.PasswordConfirmation == "" || request.Password != request.PasswordConfirmation {
		jsonFailResponse(w, http.StatusUnprocessableEntity, errors.New("password confirmation is required. must be same as password"))
		return
	}
	if request.Token == "" {
		jsonFailResponse(w, http.StatusUnprocessableEntity, errors.New("token is required"))
		return
	}
	if err := h.AuthService.VerifyAnonToken(request.Token); err != nil {
		jsonFailResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	// action
	user := &supermarket.User{
		Username: request.Username,
		Password: request.Password,
	}

	err := h.UserService.CreateUser(user)
	if err != nil {
		if h.DbErrorService.IsDuplicateError(err) {
			jsonFailResponse(w, http.StatusBadRequest, errors.New("cannot create user"))
		} else {
			jsonFailResponse(w, http.StatusBadRequest, errors.New("cannot create user"))
		}
		return
	}

	jsonSuccessResponseMessage(w, "user created")
}

func (h *Handler) Token(w http.ResponseWriter, r *http.Request) {
	token, err := h.AuthService.CreateAnonToken()
	if err != nil {
		log.Println(err)
		jsonFailResponse(w, http.StatusInternalServerError, supermarket.ErrInternalServerError)
		return
	}

	jsonSuccessResponse(w, struct {
		Token string `json:"token"`
	}{
		Token: token,
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
