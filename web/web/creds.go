package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mrtyunjaygr8/passwd/ent"
	"github.com/mrtyunjaygr8/passwd/utils"
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

type updateCredRequest struct {
	Password string `json:"password"`
}

type updateCredResponse struct {
	Body string `json:"body"`
}

func (web *Web) updateCred(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	vars := mux.Vars(r)
	var update_cred_data updateCredRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&update_cred_data)
	if err != nil {
		log.Println(err)
	}

	_, err = web.UpdateCred(token, vars["name"], update_cred_data.Password)
	if err != nil {
		log.Println(err)
		handleError(err, w)
		return
	}

	resp := updateCredResponse{Body: fmt.Sprintf("%s updated successfully", vars["name"])}
	handleJsonResponse(resp, w)
	return
}

type getCredResponse struct {
	Password string     `json:"password"`
	Cred     *ent.Creds `json:"cred"`
}

func (web *Web) getCred(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	vars := mux.Vars(r)

	cred, passwd, err := web.GetCred(token, vars["name"])
	if err != nil {
		log.Println(err)
		handleError(err, w)
		return
	}

	resp := getCredResponse{Password: passwd, Cred: cred}
	handleJsonResponse(resp, w)
	return
}

type historyPass struct {
	ID         int       `json:"id"`
	Password   string    `json:"password"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
type historyCredReponse struct {
	Data []historyPass `json:"data"`
}

func (web *Web) historyCred(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	vars := mux.Vars(r)

	history, err := web.HistoryCreds(token, vars["name"])
	if err != nil {
		log.Println(err)
		handleError(err, w)
		return
	}

	var resp_hist []historyPass
	for _, x := range history {
		resp_hist = append(resp_hist, historyPass{ID: x.ID, Password: x.Password, CreateTime: x.CreateTime, UpdateTime: x.UpdateTime})
	}

	resp := historyCredReponse{Data: resp_hist}
	handleJsonResponse(resp, w)
	return
}

type generateCredResponse struct {
	Password string     `json:"password"`
	Cred     *ent.Creds `json:"cred"`
}

func (web *Web) generateCred(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	vars := mux.Vars(r)

	str_length := r.URL.Query().Get("length")
	length := utils.DEFAULT_PASS_LENGTH
	var err error
	if str_length != "" {
		length, err = strconv.Atoi(str_length)
		if err != nil {
			log.Println(err)
			handleError(err, w)
			return
		}
	}

	str_disable_specials := r.URL.Query().Get("disable_special")
	disable_special := utils.DEFAULT_DISABLE_SPECIAL
	if str_disable_specials != "" {
		disable_special, err = strconv.ParseBool(str_disable_specials)
		if err != nil {
			log.Println(err)
			handleError(err, w)
			return
		}
	}

	cred, passwd, err := web.GeneratePassForCreds(token, vars["name"], disable_special, length)
	if err != nil {
		log.Println(err)
		handleError(err, w)
		return
	}

	resp := getCredResponse{Password: passwd, Cred: cred}
	handleJsonResponse(resp, w)
	return
}
