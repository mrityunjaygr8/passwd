package cmd

import (
	"errors"
	"log"
	"os"

	"github.com/mrtyunjaygr8/passwd/utils"
)

func getToken() string {
	var token []byte
	if _, err := os.Stat(utils.LOGIN_FILE); err == nil {
		token, err = os.ReadFile(utils.LOGIN_FILE)
		if err != nil {
			log.Fatal("an error has occurred while authenticating")
		}
	} else if errors.Is(err, os.ErrNotExist) {
		log.Fatal("You are not logged in.\nPlease log in via the log command")
	} else {
		log.Fatal(err)
	}

	return string(token)
}
