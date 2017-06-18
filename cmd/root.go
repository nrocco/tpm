package cmd

import (
    "fmt"

    homedir "github.com/mitchellh/go-homedir"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"

    "github.com/nrocco/tpm/pkg/client"
)

var cfgFile string

var TpmClient client.TpmClient

var RootCmd = &cobra.Command{
    Use:   "tpm",
    Short: "A Team Password Manager CLI Application",
    Long: ``,
    PersistentPreRun: func(cmd *cobra.Command, args []string) {
        TpmClient = client.New(
            viper.GetString("server"),
            viper.GetString("username"),
            viper.GetString("password"),
        )
    },
}

func Execute() {
    RootCmd.Execute()
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
