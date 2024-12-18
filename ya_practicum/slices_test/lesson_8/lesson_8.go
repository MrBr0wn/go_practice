package main

import "fmt"

func main() {
	slice1 := make([]int, 100) // src slice
	fmt.Println("Source slice1: ", slice1, len(slice1), cap(slice1))

	slice2 := append(slice1[:10])           // create new slice with first 10 items of slice1
	slice2 = append(slice2, slice1[90:]...) // and append last 10 items of slice1

	fmt.Println("slice1 after the slice2 definition: ", slice1, len(slice1), cap(slice1))

	fmt.Println("slice2: ", slice2, len(slice2), cap(slice2))

	// upd slice2 to show the difference
	slice2[0] = 1
	slice2[len(slice2)-1] = len(slice2)

	fmt.Println("\n\nupdated slice2: ", slice2)

	fmt.Println("\n\nreverted slice2: ")

	// reverse output
	for i := len(slice2) - 1; i >= 0; i-- {
		fmt.Println(slice2[i])
	}
}
