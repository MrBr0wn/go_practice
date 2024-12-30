package main

import "fmt"

type figures int

const (
	square figures = iota
	circle
	triangle
	thor
)

func area(f figures) (func(float64) float64, bool) {
	switch f {
	case square:
		return func(x float64) float64 { return x * x }, true
	case circle:
		return func(x float64) float64 { return 3.14159 * x * x }, true
	case triangle:
		return func(x float64) float64 { return 0.433 * x * x }, true
	default:
		return nil, false
	}
}

func main() {
	ar, ok := area(circle)
	if !ok {
		fmt.Println("Error")
		return
	}
	fmt.Println(ar(3))

	ar1, ok1 := area(thor)
	if !ok1 {
		fmt.Println("Error")
		return
	}
	fmt.Println(ar1(5))
}
