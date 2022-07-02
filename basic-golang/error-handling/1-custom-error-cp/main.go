package main

import (
	"errors"
	"fmt"
)

type ErrorInvalidData struct {
	message string
	errCode int
}

func (e ErrorInvalidData) Error() string {
	return fmt.Sprintf("error %d : %s", e.errCode, e.message)
}

func GetAge(data map[string]int, name string) (int, error) {
	if _, ok := data[name]; !ok {
		return 0, errors.New("Data not found")
	}

	if data[name] < 0 {
		// Isilah baris ini dengan return 0 dan custom error yang telah dibuat dengan message error invalid data dan errCode 500
		// TODO: answer here
		return 0, ErrorInvalidData{
			message: fmt.Sprintf("%s error invalid data", name),
			errCode: 500,
		}
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