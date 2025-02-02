package main

import "fmt"

type Stringer interface {
	String() string
}

type MyMap map[string]string

func (m MyMap) String() string {
	return fmt.Sprintf("%v", m)
}

type MyInt int

func (m MyInt) String() string {
	return fmt.Sprintf("%v", m)
}

func main() {
	var mm MyMap
	var mi MyInt

	mm = MyMap{}
	mi = MyInt(5)

	fmt.Println(mm)
	fmt.Println(mi)
	// fmt.Println(mm == mi) // False
	// fmt.Println(mm == mm) // Panic
}
