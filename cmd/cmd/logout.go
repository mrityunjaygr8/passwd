package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mrtyunjaygr8/passwd/utils"
	"github.com/spf13/cobra"
)

var logout_cmd = &cobra.Command{
	Use:   "logout",
	Short: "Log Out",
	Long:  "Log out the currently logged in user",
	Run: func(cmd *cobra.Command, args []string) {
		token := getToken()
		if token == "" {
			log.Fatal("You are not logged in")
		}
		err := os.Remove(utils.LOGIN_FILE)
		if err != nil {
			log.Fatal("An error occurred while logging you out")
		}
		fmt.Println("You have been successfully logged out")
	},
}

func init() {
	root_cmd.AddCommand(logout_cmd)
}
