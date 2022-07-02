package main

import "fmt"

func main() {
	// mengembalikan string selamat sore dengan anonymous function
	goodAfternoon := func() string {
		return "Selamat sore"
	}()

	fmt.Println(goodAfternoon)
}
