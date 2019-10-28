package main

import (
	"io/ioutil"
	"os/exec"
)

func main() {

	root := "/home/caleb-powell/git"

	files, err := ioutil.ReadDir(root)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		cmd := exec.Command("git", "push")
		cmd.Dir = root + "/" + file.Name()
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
