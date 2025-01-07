package cipherer

import (
	"errors"
	"flag"
)

var cipherMode = flag.Bool("cipher", false, "Enable cipher mode.")
var decipherMode = flag.Bool("decipher", false, "Enable decipher mode.")

func main() {
	flag.Parse()

	if *cipherMode && *decipherMode {
		utils.HaltOnerr(errors.New("please specify only one mode at a time"))
	}

	if *decipherMode {

	} else if *cipherMode {

	} else {
		utils.HaltOnErr(errors.New("unknown mode"))
	}
}
