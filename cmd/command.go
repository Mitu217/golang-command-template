package cmd

import (
	"fmt"

	"github.com/Mitu217/go-cmd-template/lib"
	"github.com/spf13/cobra"
)

const (
	CommandName    = "command"
	CommandVersion = "0.0.1"
)

type Command struct {
	logger *lib.Logger
	cmd    *cobra.Command
}

func newRootCommand(name string) *cobra.Command {
	return &cobra.Command{
		Use:  fmt.Sprintf(`%s`, name),
		Long: fmt.Sprintf(`%s is a tool for ...`, name),
	}
}

func NewCommand() *Command {
	var (
		flagVerbose bool
	)

	cmd := newRootCommand(CommandName)

	cmd.AddCommand(
		newVersionCommand(CommandName, CommandVersion),

		newHelloCommand(),
	)

	cmd.PersistentFlags().BoolVarP(&flagVerbose, "verbose", "v", false, "enable verbose log")

	return &Command{
		cmd:    cmd,
		logger: new(lib.Logger),
	}
}

func (c Command) Execute() error {
	return c.cmd.Execute()
}
