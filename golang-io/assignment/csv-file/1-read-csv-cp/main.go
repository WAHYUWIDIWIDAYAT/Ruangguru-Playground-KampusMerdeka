package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fmt.Print("dummy commit")
}

func CSVToMap(data map[string]string, fileName string) (map[string]string, error) {
	csvFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, eachRecord := range rawCSVdata {
		data[eachRecord[0]] = eachRecord[1]
	}
	return data, nil
}
