package main

import "fmt"

type IntSlice []int

func (s *IntSlice) Add(v int) {
	fmt.Printf("s: %v\n", s)
	fmt.Println("*s: ", *s)

	*s = append(*s, v)
}

func main() {
	s := make(IntSlice, 0)
	s.Add(1)
	s.Add(2)
	fmt.Println("\n", s)
}
