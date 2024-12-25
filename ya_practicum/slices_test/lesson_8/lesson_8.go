package main

import "fmt"

func main() {
	// Создайте слайс и заполните его числами от 1 до 100. 
	// Оставьте в слайсе первые и последние 10 элементов и 
	// разверните слайс в обратном порядке. 
	slice1 := make([]int, 100) // src slice
	for i := range slice1 {
		slice1[i] = i + 1
	}
	fmt.Println("Source slice1: ", slice1, len(slice1), cap(slice1))

	// slice2 := make([]int, 20)
	slice2 := append([]int(nil), slice1[:10]...) // create new slice with first 10 items of slice1
	slice2 = append(slice2, slice1[90:]...)      // and append last 10 items of slice1

	fmt.Println()
	fmt.Println("slice1 after the slice2 definition: ", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2: ", slice2, len(slice2), cap(slice2))

	// upd slice2 to show the difference
	slice2[0] = 0
	slice2[len(slice2)-1] = len(slice2)

	fmt.Println("\n\nupdated slice2: ", slice2)
	fmt.Println("\n\nslice1: ", slice1)
	fmt.Println("\n\nslice2: ", slice2)

	fmt.Printf("\n\nreverted slice2: ")

	// reverse output
	for i := len(slice2) - 1; i >= 0; i-- {
		fmt.Printf("%v ", slice2[i])
	}
}
