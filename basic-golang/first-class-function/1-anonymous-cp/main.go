package main

import "fmt"

func main() {
	//fungsi goodMorning melakukan print "selamat pagi"
	// TODO: answer here
	goodMorning("wahyu")
	fmt.Printf("jenis variabelnya %T", goodMorning)

}
func goodMorning(name string) {
	fmt.Println("Selamat Pagi " + name)
}
