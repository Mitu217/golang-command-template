package main

import (
	"fmt"
	"os"

	"github.com/Mitu217/go-cmd-template/cmd"
)

func run() error {
	cmd := cmd.NewCommand()
	return cmd.Execute()
}
func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
