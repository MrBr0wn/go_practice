package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func PrintAllFilesWithFilterClosure(path string, filter string) {
	username, err1 := os.UserHomeDir()
	if err1 != nil {
		fmt.Println("unable to get user home directory")
	}

	var walk func(string)

	walk = func(path string) {
		files, err := os.ReadDir(filepath.Join(username, path))
		if err != nil {
			fmt.Println("unable to get list of files", err)
			return
		}

		for _, f := range files {
			filename := filepath.Join(path, f.Name())

			if strings.Contains(filename, filter) {
				fmt.Println(filename)
			}

			if f.IsDir() {
				walk(filename)
			}
		}
	}
	walk(path)
}

func main() {
	PrintAllFilesWithFilterClosure("Lessons/go/ya_practicum", "maps")
}
