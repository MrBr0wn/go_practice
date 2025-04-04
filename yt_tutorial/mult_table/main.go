package main

import (
	"fmt"
)

func main() {
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Print("x\\y\t")

	for _, a := range numbers {
		fmt.Printf("%8d", a)
	}

	fmt.Print("\n\n")

	for _, a := range numbers {
		fmt.Printf("%d\t", a)
		for _, b := range numbers {
			fmt.Printf("%8d", a*b)
		}
		fmt.Println("")
	}
}
