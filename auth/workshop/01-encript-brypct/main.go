package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	hashedStr, err := encryptToBcrypt("hello")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hashed String = %s", hashedStr)
}

// encrypt string kedalam bcrypt
func encryptToBcrypt(str string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(bytes), err
}
