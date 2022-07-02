package main

import (
	"fmt"
)

// Check Point:
// Two Dimention Area
// - Input: Select Choice
// - Output: Result Operation

// Contoh:
// Input:
// 1: Rectange Area
// 2: Rectangular Area
// 3: Triangle Area
// 4: Circle Area
// - Enter choice 1, 2, 3, or 4: 1 | 2 | 3 | 4
//   - (1) Rectange Area
//   	- Masukkan sisi : 5
//   - (2) Rectangular Area
// 		- Masukkan panjang : 5
// 		- Masukkan lebar : 10
// 	 - (3) Triangle Area
// 		- Masukkan panjang alas segitiga: 5
// 		- Masukkan tinggi segitiga: 10
// 	 - (4) Circle Area
//      - Masukkan jari-jari : 4

// Output:
// - (1) Luas Persegi adalah : 25
// - (2) Luas Persegi Panjang adalah : 50
// - (3) Luas Segitiga adalah : 25
// - (4) Luas Lingkaran adalah : 50.240000

func main() {
	var (
		choice   int32
		panjang2 int32
		lebar2   int32
		jari     float32
		tinggi3  int32
		alas3    int32
		sisi     int32
	)
	// TODO: answer here
	fmt.Println("1 : Rectangle Area")
	fmt.Println("2 : Rectangular Area")
	fmt.Println("3 : Triangle Area")
	fmt.Println("4 : Circle Area")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Scan(&sisi)
		fmt.Println("luas persegi", sisi*sisi)
	case 2:
		fmt.Scan(&panjang2)
		fmt.Scan(&lebar2)
		fmt.Println("Luas Persegi panjang", panjang2*lebar2)
	case 3:
		fmt.Scan(&alas3)
		fmt.Scan(&tinggi3)
		fmt.Println("Hasil luas segitiga ", 1/2*alas3*tinggi3)
	case 4:
		fmt.Scan(&jari)
		fmt.Println("Hasil luas lingkaran ", 3.14*jari*jari)
	}
}
