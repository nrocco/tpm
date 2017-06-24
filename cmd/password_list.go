package cmd

import (
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var passwordListCommand = &cobra.Command{
	Use:   "list",
	Short: "List passwords",
	Long: `When searching for passwords in Team Password Manager you can use special
operators that can help you refine your results. Search operators are
special words that allow you to find passwords quickly and accurately.

tag:string
	Search passwords that have a tag that matches the string.

access:string
	Search passwords that have the string in the access field.

username:string
	Search passwords that have the string in the username field.

name:string
	Search passwords that have the string in the name field`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var search string

		if len(args) > 0 {
			search = strings.Join(args, " ")
		}

		passwords, err := TpmClient.PasswordList(search)
		if err != nil {
			return err
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetColumnSeparator(" ")
		table.SetBorder(false)
		table.SetRowLine(false)
		table.Append([]string{"ID", "NAME", "ACCESS INFO", "USERNAME", "TAGS"})

		for _, password := range passwords {
			table.Append([]string{strconv.FormatInt(int64(password.ID), 10), password.Name, password.AccessInfo, password.Username, password.Tags})
		}

		table.Render()

		return nil
	},
}

func init() {
	passwordCommand.AddCommand(passwordListCommand)
}
