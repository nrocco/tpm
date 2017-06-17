package cmd

import (
    "errors"
    "html/template"
    "os"

    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "github.com/nrocco/tpm/client"
)

var fafa = `Name:       {{ .Name }}
Id:         {{ .Id }}
AccessInfo: {{ .AccessInfo }}
Username:   {{ .Username }}
Password:   {{ .Password }}
Email:      {{ .Email }}
UpdatedOn:  {{ .UpdatedOn }}
Tags:       {{ .Tags }}
`

var showCmd = &cobra.Command{
    Use:   "show",
    Short: "A brief description of your command",
    Long: ``,
    RunE: func(cmd *cobra.Command, args []string) error {
        if len(args) != 1 {
            return errors.New("You need to provide a password id")
        }

        id := args[0]

        client := client.New(
            viper.GetString("server"),
            viper.GetString("username"),
            viper.GetString("password"),
        )

        password, err := client.Get(id)
        if err != nil {
            return err
        }

        tmpl, err := template.New("fafa").Parse(fafa)
        if err != nil {
            return err
        }

        tmpl.Execute(os.Stdout, password)

        return nil
    },
}

func init() {
    RootCmd.AddCommand(showCmd)
}
