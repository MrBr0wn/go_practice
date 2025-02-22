package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"phonebook/book"
	"phonebook/logger"
	"strings"
	"time"
)

func main() {
	phoneBook := make(book.PhoneBook)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Phonebook")
	fmt.Println("Available commands: add, get, delete, update, list, exit")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)
		command := parts[0]
		args := parts[1:]

		switch command {
		// add alice=12345678
		case "add":
			handleCommand(doAdd, args, phoneBook)
		// get alice
		case "get":
			handleCommand(doGet, args, phoneBook)
		case "delete":
			handleCommand(doDelete, args, phoneBook)
		case "update":
			// update alice=12345678
			handleCommand(doUpdate, args, phoneBook)
		case "list":
			handleCommand(doList, args, phoneBook)
		case "exit":
			logger.Info("Exiting phoneBook... Bye!")
			return
		default:
			logger.Warn(errors.New("Unsupported command. Try: add, get, delete, update, list, exit"))
		}
	}
}

func handleCommand(cmd func([]string, book.PhoneBook) error, args []string, phoneBook book.PhoneBook) {
	if err := cmd(args, phoneBook); err != nil {
		logger.Warn(err, "command failed :'(")
	}
}

func doAdd(args []string, phoneBook book.PhoneBook) error {
	if len(args) < 1 {
		return errors.New("missing parameters for 'add' command. Use: add name=number")
	}

	kv := strings.SplitN(args[0], "=", 2)
	if len(kv) != 2 {
		return errors.New("invalid format. Use: add name=number")
	}

	name, number := kv[0], kv[1]
	err := phoneBook.Add(name, number)
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("Added an entry: %s -> %s\n", name, number))

	return nil
}

func doGet(args []string, phoneBook book.PhoneBook) error {
	if len(args) < 1 {
		return errors.New("missing parameters for 'get' command. Use: get name")
	}

	name := args[0]

	numberData, err := phoneBook.Get(name)
	if err != nil {
		return err
	}

	unixUpdatedAt := time.Unix(numberData.LastUpdatedAt, 0)

	logger.Info(
		fmt.Sprintf("Number for %s is %s (last updated at %s)\n",
			name,
			numberData.Number,
			unixUpdatedAt.Format("2006-01-02 15:05:04"),
		),
	)

	return nil
}

func doList(_ []string, phoneBook book.PhoneBook) error {
	if len(phoneBook) == 0 {
		return errors.New("phonebook is empty")
	} else {
		results := ""

		for name, number := range phoneBook {
			results += fmt.Sprintf("%s -> %s\n", name, number.Number)
		}

		logger.Info(results)
	}

	return nil
}

func doUpdate(args []string, phoneBook book.PhoneBook) error {
	if len(args) < 1 {
		return errors.New("missing parameters for 'update' command. Use: update name=new_number")
	}

	kv := strings.SplitN(args[0], "=", 2)
	if len(kv) != 2 {
		return errors.New("invalid format. Use: update name=new_number")
	}

	name, newNumber := kv[0], kv[1]

	err := phoneBook.Update(name, newNumber)
	if err != nil {
		return err
	}
	logger.Info(
		fmt.Sprintf("Updated an entry: %s -> %s\n", name, newNumber),
	)

	return nil
}

func doDelete(args []string, phoneBook book.PhoneBook) error {
	if len(args) < 1 {
		return errors.New("missing parameters for 'delete' command. Use: delete name")
	}
	name := args[0]

	err := phoneBook.Delete(name)
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("Deleted entry for %s\n", name))

	return nil
}
