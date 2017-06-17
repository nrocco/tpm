package cmd

import (
    "fmt"
    "runtime"
    "os"

    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "github.com/nrocco/tpm/client"
)

var version string = "2.0.0-rc1"

var versionCmd = &cobra.Command{
    Use: "version",
    Short: "Show version of the client and server",
    Long: ``,
    Run: func(cmd *cobra.Command, args []string) {
        client := client.New(
            viper.GetString("server"),
            viper.GetString("username"),
            viper.GetString("password"),
        )

        fmt.Println("Client:")
        fmt.Println("  Version:    " + version)
        fmt.Println("  OS/Arch:    " + runtime.GOOS + "/" + runtime.GOARCH)
        fmt.Println("  Shell:      " + os.Getenv("SHELL"))
        fmt.Println("  User:       " + os.Getenv("USER"))

        serverVersion, err := client.Version()
        if err != nil {
            return
        }

        fmt.Println("")
        fmt.Println("Server:")
        fmt.Println("  Url:        " + client.Server)
        fmt.Println("  Version:    " + serverVersion.Version)
        fmt.Println("  Date:       " + serverVersion.Date)
        fmt.Println("  ApiVersion: " + serverVersion.Number)
    },
}

func init() {
    RootCmd.AddCommand(versionCmd)
}
