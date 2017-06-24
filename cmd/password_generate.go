package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var passwordGenerateCommand = &cobra.Command{
	Use:   "generate",
	Short: "Generate a strong, random password",
	Long:  `Generate a strong, random password using the API`,
	RunE: func(cmd *cobra.Command, args []string) error {
		password, err := TpmClient.GeneratePassword()
		if err != nil {
			return err
		}

		fmt.Println(password.Value)

		return nil
	},
}

func init() {
	passwordCommand.AddCommand(passwordGenerateCommand)
}
