package cmd

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var passwordShowCommand = &cobra.Command{
	Use:   "show",
	Short: "Show a password",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("You need to provide a password id")
		}

		password, err := TpmClient.PasswordGet(args[0])
		if err != nil {
			return err
		}

		fmt.Println("Id:         " + password.ID)
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
	passwordCommand.AddCommand(passwordShowCommand)
}
