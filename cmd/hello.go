package cmd

import (
	"github.com/Mitu217/go-cmd-template/lib"
	"github.com/spf13/cobra"
)

func newHelloCommand() *cobra.Command {
	logger := new(lib.Logger)

	return &cobra.Command{
		Use: "hello",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Success("Hello, World!!")
		},
	}
}
