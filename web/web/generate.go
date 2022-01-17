package web

import (
	"log"
	"net/http"
	"strconv"

	"github.com/mrtyunjaygr8/passwd/app"
	"github.com/mrtyunjaygr8/passwd/utils"
)

type generateResponse struct {
	Password string `json:"password"`
}

func (web *Web) generate(w http.ResponseWriter, r *http.Request) {
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
	resp := generateResponse{Password: app.Generate(length, disable_special)}
	handleJsonResponse(resp, w)
	return
}
