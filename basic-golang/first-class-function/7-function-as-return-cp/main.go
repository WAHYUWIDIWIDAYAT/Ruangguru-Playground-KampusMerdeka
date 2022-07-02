package main

import "fmt"

func main() {
	// fungsi ini akan mengembalikan fungsi berikut
	// func(x, y int) int {
	// 	return x * y
	// }
	// TODO: answer here
	area := func(x, y int) int {
		return x * y
	}
	fmt.Println(area(3, 4))

}
