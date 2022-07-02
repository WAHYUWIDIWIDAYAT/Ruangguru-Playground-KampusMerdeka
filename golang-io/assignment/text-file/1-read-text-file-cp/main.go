package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("dummyCommit")
}

// Gunakan struct untuk menyimpan data file nya ya
type FileData struct {
	Name string
	Size int
	Data []byte
}

func ReadFile(name string) (FileData, error) {
	file, err := os.Open(name)
	if err != nil {
		return FileData{}, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return FileData{}, err
	}

	return FileData{
		Name: name,
		Size: len(data),
		Data: data,
	}, nil
}
