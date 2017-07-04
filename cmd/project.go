package cmd

import (
	"github.com/spf13/cobra"
)

var projectCommand = &cobra.Command{
	Use:   "project",
	Short: "Manage projects",
	Long:  ``,
}

func init() {
	rootCmd.AddCommand(projectCommand)
}
