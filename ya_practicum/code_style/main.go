package main

import (
	"fmt"
	"main/arrint"
)

func main() {
	a := []int{1, 2, 3, 4}
	b := []int{2, 3, 5, 6, 7}

	result := arrint.Add(a, b)

	fmt.Println(result)
}
