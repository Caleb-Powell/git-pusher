package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Caleb-Powell/git-pusher/command"
	"github.com/Caleb-Powell/git-pusher/config"
)

var (
	conf = config.Config{}
	cmds = command.GitCommand{}
)

func init() {
	conf = config.GetConfig()

	cmds = command.GitCommand{
		Commands: getCommands(),
	}
}

func main() {
	if !conf.Nested {
		cmds.Run(conf.RepoPath)
		return
	}

	files, err := os.ReadDir(conf.RepoPath)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		cmds.Run(conf.RepoPath + "/" + file.Name())
	}
}

func getCommands() []*exec.Cmd {
	cmds := []*exec.Cmd{}
	cmtMsg := fmt.Sprintf("\"%s\"", conf.CommitMsg)

	commitCmd := exec.Command("git", "commit", "-a", "-m", cmtMsg)

	cmds = append(cmds, commitCmd)

	if conf.Push {
		pshCmd := exec.Command("git", "push")
		cmds = append(cmds, pshCmd)
	}
	return cmds
}
