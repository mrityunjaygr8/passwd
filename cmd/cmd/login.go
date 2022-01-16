package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/mrtyunjaygr8/passwd/app"
	"github.com/mrtyunjaygr8/passwd/utils"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var login_cmd = &cobra.Command{
	Use:   "login",
	Short: "Login and open the vault",
	Long:  "Login to the password manager and open the vault to access your passwords",
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.GetConfig()

		app := app.CreateApp(config)
		defer app.Client.Close()

		fmt.Print("Enter the password for the user: ")
		login_bytes_pass, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Println("error in reading user password")
		}

		login_password = string(login_bytes_pass)
		fmt.Println()
		token, err := app.LoginUser(login_email, login_password)
		if err != nil {
			log.Println(err)
		}

		if err == nil {
			if _, err := os.Stat(utils.LOGIN_FILE); err == nil {
				err := os.WriteFile(utils.LOGIN_FILE, []byte(token), 0644)
				if err != nil {
					log.Println("an error has occurred while loggin you in")
				}
			} else if errors.Is(err, os.ErrNotExist) {
				f, err := os.Create(utils.LOGIN_FILE)
				if err != nil {
					log.Println("an error has occurred while loggin you in")
				}
				defer f.Close()

				_, err = f.WriteString(token)
				f.Sync()
			} else {
				fmt.Println(err)
			}

			fmt.Printf("User: %v logged in successfully\n", login_email)
		}
	},
}

var login_email string
var login_password string

func init() {
	login_cmd.Flags().StringVarP(&login_email, "email", "e", "", "Email for the account to log into")
	login_cmd.MarkFlagRequired("email")

	root_cmd.AddCommand(login_cmd)
}
