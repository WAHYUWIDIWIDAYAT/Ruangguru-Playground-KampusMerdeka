package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Hello World")
}

func ScanToArray(arr *[]string, fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		return err
	}
	for _, eachRecord := range rawCSVdata {
		*arr = append(*arr, eachRecord[0])
	}
	return nil
}

func ScanToMap(dataMap map[string]string, fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		return err
	}
	for _, eachRecord := range rawCSVdata {
		dataMap[eachRecord[0]] = eachRecord[1]
	}
	return nil

}
