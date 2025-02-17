package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Errors []error

func (errs Errors) Error() string {
	var res strings.Builder
	for _, err := range errs {
		res.WriteString(err.Error())
	}

	return strings.TrimRight(res.String(), ";")
}

func MyCheck(input string) error {
	var err Errors
	if len(input) > 20 {
		err = append(err, errors.New("String should be less than 20 characters; "))
	}
	if containsDigits(input) {
		err = append(err, errors.New("String should not contain any digits; "))
	}
	if containsSpices(input) {
		err = append(err, errors.New("String should not contain more than 2 spaces; "))
	}

	if len(err) == 0 {
		return nil
	}

	return err
}

func containsDigits(s string) bool {
	for _, char := range s {
		if unicode.IsDigit(char) {
			return true
		}
	}

	return false
}

func containsSpices(s string) bool {
	sum := 0
	for _, char := range s {
		if unicode.IsSpace(char) {
			sum++
		}
	}

	return sum > 2
}

func main() {
	for {
		fmt.Printf("Enter a string (q for exit): ")
		reader := bufio.NewReader(os.Stdin)
		ret, err := reader.ReadString('\n')

		fmt.Println(err)

		if err != nil {
			fmt.Println(err)
			continue
		}
		ret = strings.TrimRight(ret, "\n")
		if ret == "q" {
			break
		}

		if err = MyCheck(ret); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("String is valid")
		}
	}
}
