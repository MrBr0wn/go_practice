package main

import (
	"bufio"
	"cipherer/processor"
	"cipherer/utils"
	"errors"
	"flag"
	"os"
)

var cipherMode = flag.Bool("cipher", false, "Enable cipher mode.")
var decipherMode = flag.Bool("decipher", false, "Enable decipher mode.")

func main() {
	flag.Parse()

	if *cipherMode && *decipherMode {
		utils.HaltOnErr(errors.New("please specify only one mode at a time"))
	}

	if *decipherMode {
		decipheredBytes, err := processor.Decrypt()
		utils.HaltOnErr(err)

		println(string(decipheredBytes))
	} else if *cipherMode {
		reader := bufio.NewReader(os.Stdin)

		print("Enter full sentence: ")

		sentence, _ := reader.ReadString('\n')

		err := processor.Encrypt(sentence)
		utils.HaltOnErr(err)
	} else {
		utils.HaltOnErr(errors.New("unknown mode"))
	}
}
