package cmd

import (
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var projectListCommand = &cobra.Command{
	Use:   "list",
	Short: "List projects",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		var search string

		if len(args) > 0 {
			search = strings.Join(args, " ")
		}

		projects, err := TpmClient.ProjectList(search)
		if err != nil {
			return err
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetColumnSeparator(" ")
		table.SetBorder(false)
		table.SetRowLine(false)
		table.Append([]string{"Id", "Name", "Updated", "Tags"})

		for _, project := range projects {
			table.Append([]string{strconv.FormatInt(int64(project.ID), 10), project.Name, project.UpdatedOn, project.Tags})
		}

		table.Render()

		return nil
	},
}

func init() {
	projectCommand.AddCommand(projectListCommand)
}
