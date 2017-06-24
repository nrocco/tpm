package cmd

import (
	"github.com/spf13/cobra"
)

var passwordCommand = &cobra.Command{
	Use:   "password",
	Short: "Manage passwords",
	Long:  ``,
}

func init() {
	RootCmd.AddCommand(passwordCommand)
}
