package cmd

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/atotto/clipboard"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/mrtyunjaygr8/passwd/app"
	"github.com/mrtyunjaygr8/passwd/utils"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var creds_cmd = &cobra.Command{
	Use:   "creds",
	Short: "Info for the saved creds",
	Long:  "Information regarding the creds saved by the user in the password manager",
}

var list_creds_cmd = &cobra.Command{
	Use:   "list",
	Short: "List the saved creds",
	Long:  "List the creds saved by the user in the password manager",
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.GetConfig()
		app := app.CreateApp(config)
		defer app.Client.Close()
		token := getToken()
		if token == "" {
			log.Fatal("You are not logged in")
		}

		creds := app.ListCreds(token)
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Name", "URL", "Username", "Last Updated"})
		for _, x := range creds {
			t.AppendRow([]interface{}{x.Name, x.URL, x.Username, x.UpdateTime})
		}

		t.Render()
	},
}

var create_cred_url, create_cred_username, create_cred_password, create_cred_name string
var create_cred_cmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new credential",
	Long:  "Create a new credential and save it in the password manager",
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.GetConfig()
		app := app.CreateApp(config)
		defer app.Client.Close()
		token := getToken()
		if token == "" {
			log.Fatal("You are not logged in")
		}
		fmt.Printf("Enter the password for %v: ", create_cred_name)
		create_creds_bytes_pass, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Println("error in reading user password")
		}

		create_cred_password = string(create_creds_bytes_pass)
		fmt.Println()

		cred, err := app.CreateCreds(token, create_cred_name, create_cred_username, create_cred_password, create_cred_url)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(cred)
	},
}

var get_cred_show bool
var get_cred_cmd = &cobra.Command{
	Use:       "get name [flags] ",
	Short:     "Get a cred",
	Long:      "Get details about a saved credential",
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"name"},
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.GetConfig()
		app := app.CreateApp(config)
		defer app.Client.Close()
		token := getToken()
		if token == "" {
			log.Fatal("You are not logged in")
		}
		cred, pass, err := app.GetCred(token, args[0])
		if err != nil {
			log.Fatal(err)
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		if get_cred_show {
			t.AppendHeader(table.Row{"Name", "URL", "Username", "Last Updated", "Password"})
			t.AppendRow(table.Row{cred.Name, cred.URL, cred.Username, cred.UpdateTime, pass})
		} else {
			t.AppendHeader(table.Row{"Name", "URL", "Username", "Last Updated"})
			t.AppendRow(table.Row{cred.Name, cred.URL, cred.Username, cred.UpdateTime})
			clipboard.WriteAll(pass)
		}

		t.Render()
	},
}

var delete_cred = &cobra.Command{
	Short:     "Delete a cred",
	Long:      "Delete details about a saved credential",
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"name"},
	Use:       "delete",
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.GetConfig()
		app := app.CreateApp(config)
		defer app.Client.Close()
		token := getToken()
		if token == "" {
			log.Fatal("You are not logged in")
		}

		err := app.DeleteCred(token, args[0])
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Cred: %v deleted successfully", args[0])
	},
}

var update_cred_password string
var update_cred = &cobra.Command{
	Short:     "Update a cred",
	Long:      "Update details about a saved credential",
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"name"},
	Use:       "update",
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.GetConfig()
		app := app.CreateApp(config)
		defer app.Client.Close()
		token := getToken()
		if token == "" {
			log.Fatal("You are not logged in")
		}

		fmt.Printf("Enter the password for %v: ", args[0])
		update_creds_bytes_pass, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Println("error in reading user password")
		}

		update_cred_password = string(update_creds_bytes_pass)
		fmt.Println()

		_, err = app.UpdateCred(token, args[0], update_cred_password)

		log.Printf("Cred: %v deleted successfully", args[0])
	},
}

var history_cmd = &cobra.Command{
	Use:       "history",
	Short:     "History of cred",
	Long:      "History of a saved credential",
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"name"},
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.GetConfig()
		app := app.CreateApp(config)
		defer app.Client.Close()
		token := getToken()
		if token == "" {
			log.Fatal("You are not logged in")
		}

		passwds, err := app.HistoryCreds(token, args[0])
		if err != nil {
			log.Fatalf("error getting cred: %v history", args[0])
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Created", "Password"})
		for _, x := range passwds {
			t.AppendRow([]interface{}{x.CreateTime, x.Password})
		}

		t.Render()
	},
}

var generate_cred_length int
var generate_cred_specials bool
var generate_cred_cmd = &cobra.Command{
	Short:     "Generate password for cred",
	Long:      "Generate password and update a saved credential",
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"name"},
	Use:       "generate",
	Run: func(cmd *cobra.Command, args []string) {
		config := utils.GetConfig()
		app := app.CreateApp(config)
		defer app.Client.Close()
		token := getToken()
		if token == "" {
			log.Fatal("You are not logged in")
		}

		_, pass, err := app.GeneratePassForCreds(token, args[0], generate_cred_specials, generate_cred_length)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Password for Cred: %v successfully updated and copied to clipboard", args[0])
		clipboard.WriteAll(pass)
	},
}

func init() {
	// Create Cred Command
	create_cred_cmd.Flags().StringVarP(&create_cred_name, "name", "n", "", "[REQUIRED]The name of the credential to be added")
	create_cred_cmd.MarkFlagRequired("name")
	create_cred_cmd.Flags().StringVarP(&create_cred_username, "username", "u", "", "[REQUIRED]The username of the credential to be added")
	create_cred_cmd.MarkFlagRequired("username")
	create_cred_cmd.Flags().StringVarP(&create_cred_url, "url", "U", "", "[REQUIRED]The URL of the credential to be added")
	create_cred_cmd.MarkFlagRequired("name")
	creds_cmd.AddCommand(create_cred_cmd)

	// Get Cred Command
	get_cred_cmd.Flags().BoolVarP(&get_cred_show, "show", "s", false, "Show the password in the Output [NOT RECOMMENDED]")
	creds_cmd.AddCommand(get_cred_cmd)

	// List Creds Command
	creds_cmd.AddCommand(list_creds_cmd)

	// Update Cred Command
	creds_cmd.AddCommand(update_cred)

	// History of Cred Command
	creds_cmd.AddCommand(history_cmd)

	// Delete Cred Command
	creds_cmd.AddCommand(delete_cred)

	// Generate Cred Command
	generate_cred_cmd.Flags().IntVarP(&generate_cred_length, "length", "l", utils.DEFAULT_PASS_LENGTH, "The length of the password to be generated")
	generate_cred_cmd.Flags().BoolVarP(&generate_cred_specials, "disable-special", "d", utils.DEFAULT_DISABLE_SPECIAL, "Disable use of Special Charecters in the password generation")
	creds_cmd.AddCommand(generate_cred_cmd)
	root_cmd.AddCommand(creds_cmd)
}
