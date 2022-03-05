package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func hashWithSalt(text string) (string, string) {
	var salt = fmt.Sprintf("%d", time.Now().UnixNano())
	var saltedText = fmt.Sprintf("%s%s", text, salt)

	var sha = sha1.New()
	sha.Write([]byte(saltedText))
	var encrypted = sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted), salt
}

func main() {
	var password = "1password1"

	var sha = sha1.New()
	sha.Write([]byte(password))
	var encrypted = sha.Sum(nil)
	var encryptedPassword = fmt.Sprintf("%x", encrypted)

	fmt.Println("Hash Password:", encryptedPassword)

	// Salting Hash-SHA1
	var salts []string
	var text = "password"
	fmt.Println("text:", text)

	for i := 0; i <= 4; i++ {
		var hash, salt = hashWithSalt(text)

		fmt.Printf("Hash %d: %s,  salt %d: %s", i+1, hash, i+1, salt)
		fmt.Println()

		salts = append(salts, salt)

		time.Sleep(time.Second * 2)
	}

	fmt.Println(salts)
}
