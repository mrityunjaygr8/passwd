package cmd

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/mrtyunjaygr8/passwd/app"
	"github.com/mrtyunjaygr8/passwd/utils"
	"github.com/spf13/cobra"
)

var generate_length int
var generate_specials bool
var generate_cmd = &cobra.Command{
	Use:   "generate",
	Short: "generate password",
	Long:  "command to generate a new password",
	Run: func(cmd *cobra.Command, args []string) {
		pass := app.Generate(generate_length, generate_specials)
		clipboard.WriteAll(pass)

		fmt.Println("A new password has been generated and copied to the clipboard")
	},
}

func init() {
	generate_cmd.Flags().IntVarP(&generate_length, "length", "l", utils.DEFAULT_PASS_LENGTH, "The length of the password to be generated")
	generate_cmd.Flags().BoolVarP(&generate_specials, "disable-special", "d", utils.DEFAULT_DISABLE_SPECIAL, "Disable use of Special Charecters in the password generation")
	root_cmd.AddCommand(generate_cmd)
}
