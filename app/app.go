package app

import (
	"context"
	"fmt"
	"log"

	"github.com/mrtyunjaygr8/passwd/ent"
	"github.com/mrtyunjaygr8/passwd/utils"
)

type App struct {
	Client  ent.Client
	Context context.Context
}

func CreateApp(config utils.Config) App {
	config_str := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_NAME, config.DB_PASS)
	client, err := ent.Open("postgres", config_str)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	app := App{}
	app.Client = *client
	app.Context = ctx

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal("failed creating schema resources: %w", err)
	}

	return app
}
