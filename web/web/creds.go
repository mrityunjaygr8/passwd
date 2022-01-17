package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrtyunjaygr8/passwd/ent"
)

type listCredResponse struct {
	Data []*ent.Creds `json:"data"`
}

func (web *Web) listCreds(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	list := web.ListCreds(token)
	resp := listCredResponse{Data: list}
	handleJsonResponse(resp, w)
	return
}

type createCredRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
	URL      string `json:"url"`
}

type createCredResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	URL      string `json:"url"`
}

func (web *Web) createCred(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	var create_cred_data createCredRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&create_cred_data)
	if err != nil {
		log.Println(err)
	}

	cred, err := web.CreateCreds(token, create_cred_data.Name, create_cred_data.Username, create_cred_data.Password, create_cred_data.URL)
	if err != nil {
		log.Println(err)
		handleError(err, w)
		return
	}

	resp := createCredResponse{Name: cred.Name, Username: cred.Username, URL: cred.URL, ID: cred.ID}
	handleJsonResponse(resp, w)
	return
}

type deleteCredResponse struct {
	Body string `json:"body"`
}

func (web *Web) deleteCred(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	vars := mux.Vars(r)
	err := web.DeleteCred(token, vars["name"])
	if err != nil {
		log.Println(err)
		handleError(err, w)
		return
	}

	resp := deleteCredResponse{Body: fmt.Sprintf("%s deleted successfully", vars["name"])}
	handleJsonResponse(resp, w)
	return
}
