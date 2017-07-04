package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/nrocco/tpm/pkg/client"
)

var tpmClient client.TpmClient

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version of the client and server",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		tpmClient = client.New(
			viper.GetString("server"),
			viper.GetString("username"),
			viper.GetString("password"),
		)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Client:")
		fmt.Println("  Version:    " + Version)
		fmt.Println("  OS/Arch:    " + runtime.GOOS + "/" + runtime.GOARCH)
		fmt.Println("  Shell:      " + os.Getenv("SHELL"))
		fmt.Println("  User:       " + os.Getenv("USER"))
		fmt.Println("")
		fmt.Println("Server:")
		fmt.Println("  Url:        " + tpmClient.Server)

		serverVersion, err := tpmClient.Version()
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
