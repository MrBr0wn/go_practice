package main

import "fmt"

func main() {
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}
	fmt.Println("Source arr: ", input)
	fmt.Println("Delete duplicates...")
	fmt.Println("Result:")
	result := removeDuplicates(input)
	fmt.Println(result)
}

// Напишите функцию, которая убирает дубликаты, сохраняя порядок слайса
func removeDuplicates(input []string) []string {
	uniq := make(map[string]int)
	result := make([]string, 0, len(input))
	for _, animal := range input {
		uniq[animal]++
	}

	for k := range uniq {
		result = append(result, k)
	}
	return result
}
