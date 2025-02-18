package main

import (
	"fmt"
	"os"
)

func myfunc() {
	if _, err := os.ReadFile("test.txt"); err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Start")
	myfunc()
	fmt.Println("Finish")
}
