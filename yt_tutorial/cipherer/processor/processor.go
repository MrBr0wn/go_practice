package processor

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"

	"golang.org/x/crypto/argon2"
	"golang.org/x/term"
)

type EncryptedPackage struct {
	Nonce         []byte
	Salt          []byte
	EncryptedData []byte
}

func DeriveKeyFrom(passphrase, salt []byte) ([]byte, error) {
	key := argon2.IDKey(passphrase, salt, 2, 64*512, 2, 32)

	return key, nil
}

func MakeCrypterFrom(key []byte) (cipher.AEAD, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	return cipher.NewGCM(block)
}

func GetPassphrase() ([]byte, error) {
	println("Enter passphrase:")
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))

	if err != nil {
		return nil, fmt.Errorf("Failed to grab passphrase: %w", err)
	}

	passphrase, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return nil, fmt.Errorf("Failed to grab passphrase: %w", err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	return passphrase, nil
}
