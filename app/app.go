package app

import (
	"fmt"
	"log"

	"github.com/mrtyunjaygr8/passwd/ent"
	"github.com/mrtyunjaygr8/passwd/utils"
)

func CreateApp(config utils.Config) ent.Client {
	config_str := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_NAME, config.DB_PASS)
	client, err := ent.Open("postgres", config_str)
	if err != nil {
		log.Fatal(err)
	}

	return *client
}
