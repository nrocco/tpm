package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"runtime"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version of the client and server",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Client:")
		fmt.Println("  Version:    " + Version)
		fmt.Println("  OS/Arch:    " + runtime.GOOS + "/" + runtime.GOARCH)
		fmt.Println("  Shell:      " + os.Getenv("SHELL"))
		fmt.Println("  User:       " + os.Getenv("USER"))
		fmt.Println("")
		fmt.Println("Server:")
		fmt.Println("  Url:        " + TpmClient.Server)

		serverVersion, err := TpmClient.Version()
		if err != nil {
			fmt.Println("  Error:      " + err.Error())
		} else {
			fmt.Println("  Version:    " + serverVersion.Version)
			fmt.Println("  Date:       " + serverVersion.Date)
			fmt.Println("  ApiVersion: " + serverVersion.Number)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
