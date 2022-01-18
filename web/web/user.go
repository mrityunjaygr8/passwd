package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mrtyunjaygr8/passwd/ent"
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
		log.Println(err)
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
		log.Println(err)
		handleError(err, w)
		return
	}

	resp := meResponse{Email: user.Email}
	handleJsonResponse(resp, w)
	return
}

type createUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type createUserResponse struct {
	User *ent.User `json:"user"`
}

func (web *Web) createUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var create_user_data createUserRequest

	err := decoder.Decode(&create_user_data)
	if err != nil {
		log.Println(err)
	}

	user, err := web.CreateUser(create_user_data.Email, create_user_data.Password)

	if err != nil {
		log.Println(err)
		handleError(err, w)
		return
	}
	resp := createUserResponse{User: user}
	handleJsonResponse(resp, w)
	return
}
