package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("dummyCommit")
}

func AddString(fileName string, stringToAdd string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(stringToAdd)
	if err != nil {
		return err
	}
	return nil
}
