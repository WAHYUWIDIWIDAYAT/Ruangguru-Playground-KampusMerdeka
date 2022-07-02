package main

import "fmt"

// Disini kalian coba untuk membaca nilai dari slice yang diberikan.
// Lalu kalian akan menambahkan kata "Olleh" pada slice tersebut.
func main() {
	slice := []string{"Hello", "World"}
	// TODO: answer here

	for index, value := range slice {
		fmt.Println(index, value)
	}
	slice = append(slice, "Olleh")

	for index, value := range slice {
		fmt.Println(index, value)
	}
}
