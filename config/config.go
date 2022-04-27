package config

import (
	"flag"
	"log"
	"os"
)

type Config struct {
	CommitMsg string
	RepoPath  string
	Push      bool
	Nested    bool
}

var (
	conf = Config{}
)

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal("failed to get pwd", err)
	}

	flag.StringVar(&conf.RepoPath, "path", pwd, "Directory git repos")
	flag.StringVar(&conf.CommitMsg, "commitMsg", "Auto Commit", "Commit message on auto commits")
	flag.BoolVar(&conf.Push, "push", false, "Auto push enable")
	flag.BoolVar(&conf.Nested, "nested", false, "Directory containing multiple git repos")

	flag.Parse()
}

func GetConfig() Config {
	return conf
}
