package main

import (
	"fmt"
)

func mypanic() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Panic has been raised: ", p)
		}
	}()
	panic("Attention! Don't panic!")
}

func main() {
	fmt.Println("Start")
	mypanic()
	fmt.Println("Finish")
}
