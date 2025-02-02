package main

import "fmt"

func Mul(a interface{}, b int) interface{} {
	switch a1 := a.(type) {
	case string:
		a2 := ""
		for b > 0 {
			a2 += a1
			b--
		}
		return a2
	case int:
		return a1 * b
	default:
		return nil
	}
}

func main() {
	a := "str1"
	b := 2
	fmt.Println(Mul(a, b))

	c := 2
	b = 3
	fmt.Println(Mul(c, b))
}
