package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	pwd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	path := flag.String("path", pwd, "This should be the directry where you hold all of your get repos. Default is the dir that it is running in.")

	flag.Parse()

	files, err := ioutil.ReadDir(*path)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		cmd := exec.Command("git", "push")
		cmd.Dir = *path + "/" + file.Name()
		cmd.Run()
	}

	// err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
	// 	files = append(files, path)
	// 	return nil
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// for _, file := range files {
	// 	fmt.Println(file)
	// }
}
