package main

import (
	"fmt"
)

// Check Point:
// Menghitung volume tabung
// - Input: jari-jari, tinggi
// - Output: volume tabung

// Contoh:
// Input:
// - Masukkan jari-jari alas tabung: 2
// - Masukkan tinggi tabung : 20
// Output:
// - Jadi volumenya adalah : 251.200012
var (
	jari   float32
	tinggi float32
	result float32
)

func main() {
	// TODO: answer here
	fmt.Scan(&jari)
	fmt.Scan(&tinggi)
	result = jari * jari * 3.14 * tinggi
	fmt.Println(result)
}
