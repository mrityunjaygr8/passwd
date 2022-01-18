package web

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrtyunjaygr8/passwd/app"
	"github.com/mrtyunjaygr8/passwd/utils"
)

type Web struct {
	app.App
	Router *mux.Router
}

type error_response struct {
	Body string `json:"body"`
	Code int    `json:"code"`
}

func CreateWebApp(config utils.Config) Web {
	web := Web{}
	app := app.CreateApp(config)

	router := mux.NewRouter()
	web.Client = app.Client
	web.Context = app.Context

	api_router := router.PathPrefix("/api/").Subrouter()

	api_router.HandleFunc("/user/create", web.createUser).Methods("POST")
	api_router.HandleFunc("/login", web.LoginHandler).Methods("POST")
	api_router.HandleFunc("/me", web.MeHandler).Methods("GET")
	api_router.HandleFunc("/generate", web.generate).Methods("GET")

	cred_router := api_router.PathPrefix("/creds/").Subrouter()

	cred_router.HandleFunc("/", web.listCreds).Methods("GET")
	cred_router.HandleFunc("/", web.createCred).Methods("POST")

	cred_router.HandleFunc("/{name}", web.deleteCred).Methods("DELETE")
	cred_router.HandleFunc("/{name}", web.updateCred).Methods("PATCH")
	cred_router.HandleFunc("/{name}", web.getCred).Methods("GET")
	cred_router.HandleFunc("/{name}/history", web.historyCred).Methods("GET")
	cred_router.HandleFunc("/{name}/generate", web.generateCred).Methods("GET")

	web.Router = router
	return web

}

func handleError(err error, w http.ResponseWriter) {
	if errors.Is(err, utils.NOT_FOUND) {
		e := error_response{Code: http.StatusNotFound, Body: "not-found"}
		handleJsonResponse(e, w)

	} else if errors.Is(err, utils.BAD_REQUEST) {
		e := error_response{Code: http.StatusBadRequest, Body: "bad-request"}
		handleJsonResponse(e, w)

	} else {
		e := error_response{Code: http.StatusInternalServerError, Body: "wierd-error-dude"}
		handleJsonResponse(e, w)

	}
}

func handleJsonResponse(data interface{}, w http.ResponseWriter) {
	js, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
