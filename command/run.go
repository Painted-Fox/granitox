package command

import (
	"github.com/mitchellh/cli"
)

type RunCommand struct {
	Ui cli.Ui
}

func (c *RunCommand) Help() string {
	return ""
}

func (c *RunCommand) Run(_ []string) int {
	return 0
}

func (c *RunCommand) Synopsis() string {
	return "Runs your test(s)."
}
