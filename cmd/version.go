package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	version   string
	commit    string
	buildDate string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version and build information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Client:")
		fmt.Printf("  Version:    %s\n", version)
		fmt.Printf("  Commit:     %s\n", commit)
		fmt.Printf("  Build Date: %s\n", buildDate)
		fmt.Printf("  Platform:   %s (%s)\n", runtime.GOOS, runtime.GOARCH)
		fmt.Printf("  Build Info: %s (%s)\n", runtime.Version(), runtime.Compiler)
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
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
