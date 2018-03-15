package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Output shell completion code for the specified shell",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("#compdef %s\n\n", rootCmd.Name())
		generateCompletions(rootCmd, []string{})
	},
}

func generateCompletions(cmd *cobra.Command, names []string) {
	names = append(names, cmd.Name())
	fullName := strings.Join(names, "_")

	if cmd.HasSubCommands() {
		fmt.Printf("local -a %s_commands\n", fullName)
		fmt.Printf("%s_commands=(\n", fullName)
		for _, command := range cmd.Commands() {
			fmt.Printf("  '%s:%s'\n", command.Name(), strings.Replace(command.Short, "'", "", -1))
		}
		fmt.Printf(")\n\n")

		fmt.Printf("if [[ CURRENT -eq %d && $words[%d] == '%s' ]]\n", len(names)+1, len(names), cmd.Name())
		fmt.Printf("then\n")
		fmt.Printf("    _describe '%s_commands' %s_commands\n", fullName, fullName)
		fmt.Printf("fi\n\n")
	} else if cmd.BashCompletionFunction != "" {
		fmt.Printf("if [[ CURRENT -eq %d && $words[%d] == '%s' ]]\n", len(names)+1, len(names), cmd.Name())
		fmt.Printf("then\n")
		fmt.Printf("    %s\n", cmd.BashCompletionFunction)
		fmt.Printf("fi\n\n")
	}

	for _, command := range cmd.Commands() {
		generateCompletions(command, names)
	}
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
