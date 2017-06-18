package cmd

import (
    "strings"
    "errors"
    "os"
    "strconv"

    "github.com/spf13/cobra"
    "github.com/olekukonko/tablewriter"
)

var searchCmd = &cobra.Command{
    Use: "search",
    Short: "Search for passwords",
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
        if len(args) == 0 {
            return errors.New("You need to provide a search query")
        }

        search := strings.Join(args, " ")

        passwords, err := TpmClient.List(search)
        if err != nil {
            return err
        }

        table := tablewriter.NewWriter(os.Stdout)
        table.SetAlignment(tablewriter.ALIGN_LEFT)
        table.SetColumnSeparator(" ")
        table.SetBorder(false)
        table.SetRowLine(false)

        for _, password := range passwords {
            table.Append([]string{strconv.FormatInt(int64(password.Id), 10), password.Name, password.AccessInfo, password.Username, password.Tags})
        }

        table.Render()

        return nil
    },
}

func init() {
    passwordCmd.AddCommand(searchCmd)
}
