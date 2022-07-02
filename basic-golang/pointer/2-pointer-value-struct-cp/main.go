package main

import "fmt"

type Rectangle struct {
	Width  int
	Length int
}

//buat struct Rectangle dengan dua atribut yaitu Width dan Length
// tambah dua method :
// SetWidthPointer(width int) dan SetWidthValue(width int)
// SetWidthPointer(width int) untuk mengubah width dengan pointer receiver
// SetWidthValue(width int) untuk mengubah width dengan value
func (r *Rectangle) SetWidthPointer(width int) {
	r.Width = width
}
func (rect Rectangle) SetWidthValue(width int) {
	rect.Width = width
}

// TODO: answer here
func main() {
	var r Rectangle
	r.Width = 10
	r.Length = 20

	fmt.Println("sebelum melakukan set width dengan pointer", r.Width)
	r.SetWidthPointer(30)
	fmt.Println("sesudah melakukan set width dengan pointer", r.Width)
	r.SetWidthValue(70)
	fmt.Println("sesudah melakukan set width dengan value", r.Width)
}
