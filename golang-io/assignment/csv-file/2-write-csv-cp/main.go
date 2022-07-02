package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

type Menu struct {
	Name  string
	Price int
}

func WriteToCSV(fileName string, records []Menu) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range records {
		err := writer.Write([]string{record.Name, strconv.Itoa(record.Price)})
		if err != nil {
			return err
		}
	}
	return nil
}
