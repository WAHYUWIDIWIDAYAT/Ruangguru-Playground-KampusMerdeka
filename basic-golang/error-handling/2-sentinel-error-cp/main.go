package main

import (
	// import
	"errors"
	"fmt"
)

// Dari contoh yang diberikan, buatlah sentinel error tambahan untuk error handling pada function FindData
// Kalian dapat menambahkan sentinel error untuk mengecek apakah umurnya valid (misal kurang dari 0) dengan nama ErrInvalidAge

var ErrDataNotFound = errors.New("error data not found")
var ErrInvalidAge = errors.New("error age is invalid, less than 0")

// TODO: answer here

func GetAge(data map[string]int, name string) (int, error) {
	if _, ok := data[name]; !ok {
		return 0, ErrDataNotFound
	}

	if data[name] < 0 {
		ErrInvalidAge := errors.New("error age is invalid, less than 0")
		return 0, ErrInvalidAge
	}

	return data[name], nil
}

func main() {
	peopleAge := map[string]int{
		"John": 20,
		"Raz":  8,
		"Tony": -1,
	}

	_, err := GetAge(peopleAge, "Tony")
	if err != nil {
		fmt.Println(err.Error())
	}
}
