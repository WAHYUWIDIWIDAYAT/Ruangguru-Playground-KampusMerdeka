package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

//Buat struct Student dengan field Name tipe data <string> dan ScoreAverage tipe data <float64>
//tampilkan dengan wording:
//Hello Rogu,
//Nilai rata rata kamu 7.8

type Student struct {
	Name         string
	ScoreAverage float64
}

// main function
func main() {
	buff := new(bytes.Buffer)
	// TODO: answer here
	tmp1 := template.Must(template.New("tmp1").Parse(`Hello {{.Name}},Nilai rata rata kamu {{.ScoreAverage}}`))
	std := Student{
		Name:         "Rogu",
		ScoreAverage: 7.8,
	}
	if err := tmp1.Execute(buff, std); err != nil {
		log.Fatalf("execute template error: %s", err.Error())
	}
	fmt.Println(buff.String())
}
