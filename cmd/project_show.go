package cmd

import (
	"errors"
	"fmt"
	"strconv"

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

		id, _ := strconv.Atoi(args[0])

		project, err := TpmClient.ProjectGet(id)
		if err != nil {
			return err
		}

		fmt.Println("Id:         " + strconv.FormatInt(int64(project.ID), 10))
		fmt.Println("Name:       " + project.Name)
		fmt.Println("UpdatedOn:  " + project.UpdatedOn)
		fmt.Println("Tags:       " + project.Tags)

		return nil
	},
}

func init() {
	projectCommand.AddCommand(projectShowCommand)
}
