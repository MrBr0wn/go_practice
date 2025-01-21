package main

import (
	"fmt"
)

type MyStruct struct {
	A     int
	Log   func(s string)
	panik func(i int)
}

func main() {
	var s = MyStruct{
		A:   1,
		Log: func(s string) { fmt.Println(s) },
	}

	s.Log("some string")
	// s.panik(3) // will be panic
}
