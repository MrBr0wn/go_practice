package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	phoneBook := make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Phonebook")
	fmt.Println("Available commands: add, get, delete, update, list, exit")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)
		command := parts[0]

		switch command {
		// add alice=12345678
		case "add":
			kv := strings.SplitN(parts[1], "=", 2)
			if len(kv) != 2 {
				fmt.Println("Invalid format. Use: add name=number")
				continue
			}

			name, number := kv[0], kv[1]
			phoneBook[name] = number
			fmt.Printf("Added/Updated: %s -> %s\n", name, number)
		// get alice
		case "get":
			name := parts[1]
			number, exists := phoneBook[name]
			if exists {
				fmt.Printf("Number for %s is %s\n", name, number)
			} else {
				fmt.Printf("No entry found for %s\n", name)
			}
		case "delete":
			name := parts[1]
			_, exists := phoneBook[name]
			if exists {
				delete(phoneBook, name)
				fmt.Printf("Deleted entry for %s\n", name)
			} else {
				fmt.Printf("No entry to delete for %s\n", name)
			}
		case "update":
			// update alice=12345678
			kv := strings.SplitN(parts[1], "=", 2)
			if len(kv) != 2 {
				fmt.Println("Invalid format. Use: update name=number")
				continue
			}
			name, newNumber := kv[0], kv[1]
			_, exists := phoneBook[name]
			if exists {
				phoneBook[name] = newNumber
				fmt.Printf("Updated %s -> %s\n", name, newNumber)
			} else {
				fmt.Println("No such entry to update")
			}
		case "list":
			if len(phoneBook) == 0 {
				fmt.Println("Phonebook is empty")
			} else {
				for name, number := range phoneBook {
					fmt.Printf("%s -> %s\n", name, number)
				}
			}
		case "exit":
			fmt.Println("Exiting phoneBook...")
			return
		default:
			fmt.Println("Unsupported command. Try: add, get, delete, update, list, exit")
		}
	}
}
