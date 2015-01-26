package main

import (
	"strings"
	"github.com/gophergala/go_ne/plugins/shared"
	"github.com/rosstimson/go_ne/plugins/core"
)

type Command struct {
}

func (t *Command) Execute(args shared.Args, reply *shared.Response) error {
	var commands []shared.Command

	commands = append(commands, installCommand(args))

	*reply = shared.NewResponse(commands...)

	return nil
}

func NewCommand() *Command {
	return new(Command)
}

func installCommand(args shared.Args) shared.Command {
	packages := shared.ExtractOptions(args.Options["packages"])

	cmd := []string{
		"sudo",
		"yum",
		"install",
		"-y",
		strings.Join(packages, " "),
	}

	return shared.NewCommand(strings.Join(cmd, " "))
}

func main() {
	plugin.Register(NewCommand())
	plugin.Serve()
}
