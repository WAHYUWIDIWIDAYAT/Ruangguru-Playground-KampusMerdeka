package main

import "fmt"

//fungsi square sama seperti sebelumnya
//yang membedakan adalah fungsi ini
//menggunakan blank return
func main() {
	result1, result2 := square(4, 5)
	fmt.Printf("%d dan %d\n", result1, result2)
	fmt.Println(square(9, 8))
}
func square(angka1, angka2 int) (a, b int) {
	a = angka1 * angka2
	b = angka1 * angka2
	return
}

// TODO: answer here
