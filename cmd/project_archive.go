package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var projectArchiveCommand = &cobra.Command{
	Use:   "archive",
	Short: "Archive a project",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("You need to provide a project id")
		}

		err := TpmClient.ProjectArchive(args[0])
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	projectCommand.AddCommand(projectArchiveCommand)
}
