package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 3, 2, 1, 3, 4, 5, 6, 7, 8, 4, 6, 4, 9}
	k := 8
	fmt.Printf("Find numbers in %v, for %v\n\n", arr, k)
	fmt.Println(find(arr, k))
}

// Дан массив целых чисел A и целое число k. Нужно найти и вывести индексы пары чисел, 
// сумма которых равна k. Если таких чисел нет, то вернуть пустой слайс. 
// Индексы можно вернуть в любом порядке.
func find(arr []int, k int) []int {
	// create empty map
	m := make(map[int]int)
	// set index as key, number as value
	for i, a := range arr {
		// if m has k-a number then arr[j] + arr[i] = k. Solution.
		if j, ok := m[k-a]; ok {
			return []int{i, j}
		}
		// else save value to the m
		m[a] = i
	}
	// does not find any pairs of numbers
	return nil
}
