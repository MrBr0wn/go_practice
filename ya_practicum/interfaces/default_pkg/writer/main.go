package main

import (
	"fmt"
	"main/hashbyte"
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

	hasher := hashbyte.New(0)
	// hasher.Write([]byte{1, 2, 3, 4, 5})
	// fmt.Println("buf: ", buf)
	hasher.Write(buf)
	fmt.Printf("Hash: %v \n", hasher.Hash())
}
