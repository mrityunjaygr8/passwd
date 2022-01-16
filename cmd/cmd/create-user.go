package cmd

import (
	"fmt"
	"log"
	"syscall"

	"github.com/mrtyunjaygr8/passwd/app"
	"github.com/mrtyunjaygr8/passwd/utils"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var create_user_email string
var create_user_password string
var create_user_cmd = &cobra.Command{
	Use:   "create-user",
	Short: "Create a user",
	Long:  "command to create a new user for the password manager",
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.GetConfig()

		app := app.CreateApp(config)
		defer app.Client.Close()

		fmt.Print("Enter the password for the user: ")
		if create_user_password == "" {
			create_user_bytes_pass, err := term.ReadPassword(int(syscall.Stdin))
			if err != nil {
				log.Println("error in reading user password")
			}

			create_user_password = string(create_user_bytes_pass)
			fmt.Println()
		}

		newUser, err := app.CreateUser(create_user_email, create_user_password)
		if err != nil {
			log.Println("error in creating user: %w", err)
		}
		fmt.Printf("User: %v has been successfully created\n", newUser.Email)
	},
}

func init() {
	create_user_cmd.Flags().StringVarP(&create_user_email, "email", "e", "", "Email for the user to be created")
	create_user_cmd.MarkFlagRequired("email")
	create_user_cmd.Flags().StringVarP(&create_user_password, "password", "p", "", "Password for the user to be created [Not Recommended]")
	root_cmd.AddCommand(create_user_cmd)
}
