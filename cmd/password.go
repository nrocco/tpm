package cmd

import (
    "github.com/spf13/cobra"
)

var passwordCmd = &cobra.Command{
    Use:   "password",
    Short: "Manage passwords",
    Long: ``,
}

func init() {
    RootCmd.AddCommand(passwordCmd)
}
