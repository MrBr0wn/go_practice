package main

import "fmt"
import "strings"
import "math/big"

func main() {
	fmt.Println("Enter hex number or 'stop' to exit:")

	var input string
	for {
		fmt.Scanln(&input)

		if input = strings.ToLower(input); input == "stop" {
			break
		}

		i := new(big.Int)
		if _, status := i.SetString(processHex(input), 16); !status {
			fmt.Println("Converatation error")
			continue
		}

		fmt.Println("New value: ", i)
	}
}

func processHex(hexStr string) string {
	result := strings.TrimPrefix(hexStr, "0x")
	return result
}
