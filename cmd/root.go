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

    RootCmd.PersistentFlags().StringP("server", "s", "", "The base url of the Team Password Manager server")
    RootCmd.PersistentFlags().StringP("username", "u", "", "Username")
    RootCmd.PersistentFlags().StringP("password", "p", "", "Password")

    viper.BindPFlag("server", RootCmd.PersistentFlags().Lookup("server"))
    viper.BindPFlag("username", RootCmd.PersistentFlags().Lookup("username"))
    viper.BindPFlag("password", RootCmd.PersistentFlags().Lookup("password"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
    if cfgFile != "" {
        viper.SetConfigFile(cfgFile)
    } else {
        home, err := homedir.Dir()
        if err == nil {
            viper.AddConfigPath(home)
            viper.SetConfigName(".tpm")
        }
    }

    viper.SetEnvPrefix("tpm")
    viper.AutomaticEnv()

    err := viper.ReadInConfig()
    if err != nil {
        fmt.Println(err)
    }
}
