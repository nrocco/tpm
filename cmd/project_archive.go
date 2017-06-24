package cmd

import (
	"errors"
	"strconv"

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

		id, _ := strconv.Atoi(args[0])

		err := TpmClient.ProjectArchive(id)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	projectCommand.AddCommand(projectArchiveCommand)
}
