package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/SahilS1G/encryption-files/filecrypt"
	"golang.org/x/term"
)

func main() {
	fmt.Println("hello")
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}
	function := os.Args[1]

	switch function {
	case "help":
		printHelp()
	case "encrypt":
		encryptHandel()
	case "decrypt":
		decryptHandle()
	default:
		fmt.Println("Run encrypt to encrypt a file, and decrypt to decrypt the file")
		printHelp()
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("hello")
}

func encryptHandel() {
	if len(os.Args) < 3 {
		fmt.Println("missing the path to the file , for more info run the help command")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("file not found")
	}
	password := getPassword()
	fmt.Println("\nEncrypting...")
	filecrypt.Encrypt(file, password)
	fmt.Println("\n file successfully decrypted")

}

func decryptHandle() {
	if len(os.Args) < 3 {
		fmt.Println("missing the path to the file , for more info run the help command")
		os.Exit(0)
	}
}

func getPassword() []byte {
	fmt.Print("Enter password")
	password, _ := term.ReadPassword(0)
	fmt.Print("\nConfirm Password")
	password2, _ := term.ReadPassword(0)
	if !validatePassword(password, password2) {
		fmt.Print("\nPasswords dont match, please try again")
	}
	return password
}

func validatePassword(password1 []byte, password2 []byte) bool {
	if !bytes.Equal(password1, password2) {
		return false
	}
	return true
}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
