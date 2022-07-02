package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

//Buat function untuk menghitung average score siswa
//panggil function didalam template

type Student struct {
	Name   string
	Scores []float64
}

func (s Student) CalculateScore(scores []float64) float64 {
	// TODO: answer here
	sum := 0.0
	for _, score := range scores {
		sum += score
	}

	return sum / float64(len(scores))

}

func (s Student) GenerateStudentTemplate() string {
	buff := new(bytes.Buffer)
	// TODO: answer here

	// register function calculateScore
	avg := template.FuncMap{
		"calculateScore": s.CalculateScore,
	}

	// bikin tempalte
	tmp1, _ := template.New("student").Funcs(avg).Parse("Hello {{.Name}}, Nilai rata-rata kamu {{calculateScore .Scores}}")

	// eksekusi template
	if err := tmp1.Execute(buff, s); err != nil {
		log.Fatalf("execute template error: %s", err.Error())
	}

	return buff.String()
}

func NewStudent(name string, scores []float64) Student {
	return Student{name, scores}
}

// main function
func main() {
	std := NewStudent("Rogu", []float64{10, 11, 12})
	fmt.Println(std.GenerateStudentTemplate())
}
