package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	PrintAllFiles("Lessons/go/ya_practicum")
}

func PrintAllFiles(path string) {
	// get list of all elements from path
	user_name, err1 := os.UserHomeDir()
	files, err := os.ReadDir(filepath.Join(user_name, path))
	if err != nil {
		fmt.Println("unable to get list of files", err)
		return
	}

	if err1 != nil {
		fmt.Println("unable to get user home directory", err1)
		return
	}
	// print all list
	for _, f := range files {
		// get element name
		filename := filepath.Join(path, f.Name())
		fmt.Println(filename)
		if f.IsDir() {
			PrintAllFiles(filename)
		}
	}
}
