package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var projectUnarchiveCommand = &cobra.Command{
	Use:   "unarchive",
	Short: "Unarchive a project",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("You need to provide a project id")
		}

		err := TpmClient.ProjectUnarchive(args[0])
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	projectCommand.AddCommand(projectUnarchiveCommand)
}
