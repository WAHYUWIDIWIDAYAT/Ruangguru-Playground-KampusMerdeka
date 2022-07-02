package main

import "fmt"

func main() {
	//fungsi untuk menampilkan string dan memiliki parameter fungsi
	//fungsi yang akan dimasukkan adalah fungsi yang memberi selamat malam
	// TODO: answer here
	goodNight := func(greeting func() string) {
		fmt.Println(greeting())
	}
	goodNight(func() string {
		return "Selamat malam"
	})
}
