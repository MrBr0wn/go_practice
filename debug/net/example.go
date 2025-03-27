package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	// ваш код
	go Server()

	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	for i := 0; i < 2; i++ {
		nn, err1 := conn.Read(buf)
		if err1 != nil {
			fmt.Println("Reading error", err1)
			return
		}

		fmt.Print(strings.ToUpper(string(buf[:nn])))
	}
}

func Server() {

	listener, err := net.Listen("tcp", "127.0.0.1:8081")

	if err != nil {

		log.Println(err)

	}

	conn, err := listener.Accept()

	if err != nil {

		log.Println(err)

	}

	defer conn.Close()

	conn.Write([]byte("message"))

	time.Sleep(10)

	conn.Write([]byte("MesSaGe"))

	time.Sleep(10)

	conn.Write([]byte("MESSAGE"))

}
