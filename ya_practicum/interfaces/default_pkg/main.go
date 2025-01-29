package main

import (
	"fmt"
	"main/randbyte"
	"time"
)

func main() {
	generator := randbyte.New(time.Now().UnixNano())

	buf := make([]byte, 16)

	for i := 0; i < 5; i++ {
		n, _ := generator.Read(buf)
		fmt.Printf("Generate bytes: %v size(%d)\n", buf, n)
	}
}
