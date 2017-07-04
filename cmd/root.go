package cmd

import (
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/nrocco/tpm/pkg/client"
)

var (
	// Version holds the version number of the tpm cli tool
	Version string

	// cfgFile holds the location to the cli configuration file
	cfgFile string

	// TpmClient represents an instance of client.TpmClient
	TpmClient client.TpmClient
)

var rootCmd = &cobra.Command{
	Use:   "tpm",
	Short: "A Team Password Manager CLI Application",
	Long:  ``,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		TpmClient = client.New(
			viper.GetString("server"),
			viper.GetString("username"),
			viper.GetString("password"),
		)
	},
}

// Execute executes the rootCmd logic and is the main entry point for tpm
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tpm.yaml)")

	rootCmd.PersistentFlags().StringP("server", "s", "", "The base url of the Team Password Manager server")
	rootCmd.PersistentFlags().StringP("username", "u", "", "Username")
	rootCmd.PersistentFlags().StringP("password", "p", "", "Password")

	viper.BindPFlag("server", rootCmd.PersistentFlags().Lookup("server"))
	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
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

	viper.ReadInConfig()
}
