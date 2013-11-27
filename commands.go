package main

import (
	"github.com/Painted-Fox/granitox/command"
	"github.com/mitchellh/cli"
	"os"
)

// Commands is the mapping of all available Granitox commands.
var Commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{
		Writer: os.Stdout,
	}

	Commands = map[string]cli.CommandFactory {
		"run": func() (cli.Command, error) {
			return &command.RunCommand {
				Ui: ui,
			}, nil
		},
	}
}
