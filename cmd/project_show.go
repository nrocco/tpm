package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var projectShowCommand = &cobra.Command{
	Use:   "show",
	Short: "Show a single project",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("You need to provide a project id")
		}

		project, err := TpmClient.ProjectGet(args[0])
		if err != nil {
			return err
		}

		fmt.Println("Id:         " + project.ID)
		fmt.Println("Name:       " + project.Name)
		fmt.Println("UpdatedOn:  " + project.UpdatedOn)
		fmt.Println("Tags:       " + project.Tags)

		return nil
	},
}

func init() {
	projectCommand.AddCommand(projectShowCommand)
}
