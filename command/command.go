package command

import (
	"os/exec"
)

type GitCommand struct {
	Commands []*exec.Cmd
}

func (gc *GitCommand) Run(dir string) {
	for _, cmd := range gc.Commands {
		cmd.Dir = dir
		cmd.Run()
	}
}
