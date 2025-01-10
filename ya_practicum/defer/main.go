package main

import (
	"fmt"
	"time"
)

func main() {
	defer metricTime(time.Now())
	fmt.Printf("unintuitive() => %s, intuitive() => %s", unintuitive(), intuitive())
	println("")
}

func unintuitive() (value string) {
	defer func() { value = "Actually" }()

	return "It would seem"
}

func intuitive() string {
	value := "It would seem"
	defer func() { value = "Actually" }()

	return value
}

func metricTime(start time.Time) {
	fmt.Println(time.Now().Sub(start))
}
