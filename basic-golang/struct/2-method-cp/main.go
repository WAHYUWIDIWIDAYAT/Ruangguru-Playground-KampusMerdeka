package main

import "fmt"

// buat struct Rectangle (persegi panjang) dengan dua atribut yaitu Width dan Length
// tambah dua method :
// GetArea() dan GetPerimeter()
// GetArea() digunakan untuk menampilkan (print) luas persegi panjang
// GetPerimeter() digunakan untuk menampilkan (print) keliling persegi panjang
type Rectangle struct {
	Width  int
	Length int
}

func (r Rectangle) GetArea() int {
	luas := r.Width * r.Length
	return luas
}
func (r Rectangle) GetPerimeter() int {
	keliling := 2 * (r.Width + r.Length)
	return keliling
}

// TODO: answer here
func main() {
	var r Rectangle
	r.Width = 10
	r.Length = 20
	fmt.Println(r)

	fmt.Println(r.GetArea())
	fmt.Println(r.GetPerimeter())
}
