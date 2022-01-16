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
	Body string
	Code int
}

func CreateWebApp(config utils.Config) Web {
	web := Web{}
	app := app.CreateApp(config)

	router := mux.NewRouter()
	web.Client = app.Client
	web.Context = app.Context

	router.HandleFunc("/login", web.LoginHandler).Methods("POST")
	router.HandleFunc("/me", web.MeHandler).Methods("GET")

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
