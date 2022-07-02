package main

import "fmt"

// Kalian udah tau tuh cara memasukan nilai ke array.
// Coba kalian masukan nama kalian ke array.
// inisialisasi array nya menggunakan "var" ya.
// Lalu outputkan array nya.
func main() {
	var array1 [26]string
	array1[0] = "A"
	array1[1] = "B"
	array1[2] = "C"
	array1[3] = "D"
	array1[4] = "E"
	array1[5] = "F"
	array1[6] = "G"
	array1[7] = "H"
	array1[8] = "I"
	array1[9] = "J"
	array1[10] = "K"
	array1[11] = "L"
	array1[12] = "M"
	array1[13] = "N"
	array1[14] = "O"
	array1[15] = "P"
	array1[16] = "Q"
	array1[17] = "R"
	array1[18] = "S"
	array1[19] = "T"
	array1[20] = "U"
	array1[21] = "V"
	array1[22] = "W"
	array1[23] = "X"
	array1[24] = "Y"
	array1[25] = "Z"

	var array2 [26]int
	array2[0] = 65
	array2[1] = 66
	array2[2] = 67
	array2[3] = 68
	array2[4] = 69
	array2[5] = 70
	array2[6] = 71
	array2[7] = 72
	array2[8] = 73
	array2[9] = 74
	array2[10] = 75
	array2[11] = 76
	array2[12] = 77
	array2[13] = 78

	fmt.Println(array1[0], array1[25])
	fmt.Println(array2[0], array2[25])
}
