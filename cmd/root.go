package cmd

import (
    "fmt"
    "os"

    homedir "github.com/mitchellh/go-homedir"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var (
    cfgFile string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
    Use:   "tpm",
    Short: "A Team Password Manager CLI Application",
    Long: ``,
}

func Execute() {
    if err := RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
    cobra.OnInitialize(initConfig)

    RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tpm.yaml)")

    RootCmd.PersistentFlags().StringP("base-url", "s", "http://127.0.0.1", "Team Password Manager base url")
    RootCmd.PersistentFlags().StringP("username", "u", "from-defaults", "Username")
    RootCmd.PersistentFlags().StringP("password", "p", "@dm!n", "Password")

    viper.BindPFlag("base_url", RootCmd.PersistentFlags().Lookup("base-url"))
    viper.BindPFlag("username", RootCmd.PersistentFlags().Lookup("username"))
    viper.BindPFlag("password", RootCmd.PersistentFlags().Lookup("password"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
    if cfgFile != "" {
        // Use config file from the flag.
        viper.SetConfigFile(cfgFile)
    } else {
        // Find home directory.
        home, err := homedir.Dir()
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }

        // Search config in home directory with name ".tpm" (without extension).
        viper.AddConfigPath(home)
        viper.SetConfigName(".tpm")
    }

    viper.SetEnvPrefix("tpm")
    viper.AutomaticEnv() // read in environment variables that match

    // If a config file is found, read it in.
    if err := viper.ReadInConfig(); err == nil {
        // fmt.Println("Using config file:", viper.ConfigFileUsed())
    }
}
