package main

import "fmt"

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan address operator dan indirect operator.

// Print alamat memori dari masing-masing variabel dibawah ini
func main() {
	name := "Roger"
	age := 20
	isMarried := true

	addressname := &name
	addressage := &age
	addressIsmarried := &isMarried

	fmt.Printf("Alamat memory dari variabel name adalah %v\n", addressname)
	fmt.Printf("Alamat memory dari variabel age adalah %v\n", addressage)
	fmt.Printf("Alamat memory dari variabel ismarried adalah %v\n", addressIsmarried)
	// TODO: answer here

	fmt.Printf("Nilai dari alamat memory %v adalah %s\n", addressname, *addressname)
	fmt.Printf("Nilai dari alamat memory %v adalah %d\n", addressage, *addressage)
	fmt.Printf("Nilai dari alamat memory %v adalah %t\n", addressIsmarried, *addressIsmarried)
}
