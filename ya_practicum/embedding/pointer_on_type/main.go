package main

import (
	"fmt"
)

type Person struct {
	Name string
	Year int
}

type Student struct {
	*Person
	Group string
}

func ChangeName(s Student, name string) {
	s.Name = name
}

func main() {
	s := Student{&Person{Name: "Alex", Year: 1999}, "021"}
	fmt.Println(s.Name)
	ChangeName(s, "Curtis")
	fmt.Println(s.Name)
}
