package main

import "fmt"

var Global = 5

func main() {
	// var Global = 5

	fmt.Println("Source value: ", Global)

	updGlobal(Global)

	fmt.Println("Final value: ", Global)
}

func updGlobal(val int) {
	srcValue := val
	defer func() {
		Global = 10
		fmt.Println("Updated Global: ", Global)
		Global = srcValue
	}()
}
