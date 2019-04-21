package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newVersionCommand(commandName, version string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: fmt.Sprintf("Print %s version", commandName),
		Long:  fmt.Sprintf("Print %s version", commandName),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s version %s\n", commandName, version)
		},
	}
}
