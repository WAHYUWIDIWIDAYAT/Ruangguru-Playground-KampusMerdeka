package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

//jalankan kode ini di dalam directory yang sama dengan lokasi kode
//hal ini perlu dilakukan agar file yang dicari dapat ditemukan

type user struct {
	Name   string `json:"name"`
	School string `json:"school"`
	Class  string `json:"class"`
	Phone  string `json:"phone"`
	Score  int    `json:"score"`
}

func (s user) String() string {
	return fmt.Sprintf("name: %v\nschool: %v\nclass: %v\nphone: %v\nscore: %v\n", s.Name, s.School, s.Class, s.Phone, s.Score)
}

func main() {
	fileName := "data/read"
	json, err := readJSON(fileName)
	if err != nil {
		fmt.Println(err)
	}

	for _, user := range json {
		fmt.Println(user)
	}
}

func readJSON(fileName string) ([]user, error) {
	path, err := filepath.Abs(fileName + ".json")
	if err != nil {
		return nil, err
	}
	file, err := openFile(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	users := []user{}

	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &users)
	return users, nil
}

func openFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return file, nil
}
