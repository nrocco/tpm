package cmd

import (
    "fmt"
    "errors"
    "strconv"

    "github.com/spf13/cobra"
    "github.com/fatih/color"
)

var showCmd = &cobra.Command{
    Use:   "show",
    Short: "Show a single password",
    Long: ``,
    RunE: func(cmd *cobra.Command, args []string) error {
        if len(args) != 1 {
            return errors.New("You need to provide a password id")
        }

        id := args[0]

        password, err := TpmClient.Get(id)
        if err != nil {
            return err
        }

        fmt.Println("Id:         " + strconv.FormatInt(int64(password.Id), 10))
        fmt.Println("Name:       " + password.Name)
        fmt.Println("AccessInfo: " + password.AccessInfo)
        fmt.Println("Username:   " + password.Username)

        fmt.Print("Password:   ")
        red := color.New(color.FgRed)
        red.Add(color.BgRed)
        red.Println(password.Password)

        fmt.Println("Email:      " + password.Email)
        fmt.Println("UpdatedOn:  " + password.UpdatedOn)
        fmt.Println("Tags:       " + password.Tags)

        return nil
    },
}

func init() {
    passwordCmd.AddCommand(showCmd)
}
