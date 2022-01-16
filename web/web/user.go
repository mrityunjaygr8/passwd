package web

import (
	"encoding/json"
	"log"
	"net/http"
)

type loginResponse struct {
	Token string `json:"token"`
}

type loginRequest struct {
	Email    string
	Password string
}

func (web *Web) LoginHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var login_data loginRequest

	err := decoder.Decode(&login_data)
	if err != nil {
		log.Println(err)
	}

	token, err := web.LoginUser(login_data.Email, login_data.Password)
	if err != nil {
		handleError(err, w)
		return
	}

	resp := loginResponse{Token: token}
	handleJsonResponse(resp, w)
	return
}

type meResponse struct {
	Email string `json:"email"`
}

func (web *Web) MeHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	user, err := web.GetUser(token)
	if err != nil {
		handleError(err, w)
		return
	}

	resp := meResponse{Email: user.Email}
	handleJsonResponse(resp, w)
	return
}
