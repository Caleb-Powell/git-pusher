package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	root := "/home/caleb-powell/git"

	files, err := ioutil.ReadDir(root)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
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
