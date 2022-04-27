package main

import (
	"os"
	"os/exec"

	"github.com/Caleb-Powell/git-pusher/config"
	"github.com/Caleb-Powell/git-pusher/command"
)

var (
	conf = config.Config{}
	cmd = command.GitCommand{}
)

func init() {
	conf = config.GetConfig()
}

func main() {
	files, err := os.ReadDir(conf.RepoPath)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		cmd := exec.Command("git", "push")
		cmd.Dir = conf.RepoPath + "/" + file.Name()
		cmd.Run()
	}
}

func getCommands() []*exec.Cmd {
	if 
}
