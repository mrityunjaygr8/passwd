package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/mrtyunjaygr8/passwd/app"
	"github.com/mrtyunjaygr8/passwd/utils"
	"github.com/spf13/cobra"

	_ "github.com/lib/pq"
)

var RootCmd = &cobra.Command{
	Use:   "passwd",
	Short: "CLI for password management",
	Long:  `CLI for password management written in go.`,
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.GetConfig()
		fmt.Println(config)

		client := app.CreateApp(config)
		defer client.Close()

		ctx := context.Background()
		if err := client.Schema.Create(ctx); err != nil {
			log.Fatal("failed creating schema resources: %w", err)
		}

		newUser, err := client.User.Create().SetEmail("msyt119691@gmail.com").SetPassword("dr0w.Ssap").Save(ctx)
		if err != nil {
			log.Println("error in creating user: %w", err)
		}
		fmt.Println(newUser)
	},
}
