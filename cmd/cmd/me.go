package cmd

import (
	"fmt"
	"log"

	"github.com/mrtyunjaygr8/passwd/app"
	"github.com/mrtyunjaygr8/passwd/utils"
	"github.com/spf13/cobra"
)

var me_cmd = &cobra.Command{
	Use:   "me",
	Short: "Get my details",
	Long:  "Get the details of the currently logged in user",
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.GetConfig()

		app := app.CreateApp(config)
		defer app.Client.Close()
		out, err := app.GetUser(getToken())
		if err != nil {
			log.Fatal("An error has occurred: ", err)
		}

		fmt.Println("User info: ", fmt.Sprintf("%v-%v", out.ID, out.Email))
	},
}

func init() {
	root_cmd.AddCommand(me_cmd)
}
