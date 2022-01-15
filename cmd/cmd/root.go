package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	_ "github.com/lib/pq"
)

var rootCmd = &cobra.Command{
	Use:   "passwd",
	Short: "CLI for password management",
	Long:  `CLI for password management written in go.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
