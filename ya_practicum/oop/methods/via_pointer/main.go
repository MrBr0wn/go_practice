package main

import "fmt"

type MyType struct {
	value int
}

func (t *MyType) SetValue(v int) {
	t.value = v
}

func (t MyType) String() string {
	return fmt.Sprintf("Value: %d", t.value)
}

func main() {
	t := MyType{}
	t.SetValue(100)
	fmt.Println(t)

	d := &MyType{}
	d.SetValue(200)
	fmt.Println(d)

	fmt.Println(d.String())
}
