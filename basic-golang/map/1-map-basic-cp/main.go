package main

import "fmt"

// Buatlah map dengan key "Nama Provinsi" pada pulau Kalimantan dan value nya adalah "Ibu Kota" provinsi tersebut
// Output satu semua key dan value yang ada di map tersebut
func main() {
	// TODO: answer here
	kamus := map[string]string{"kalimantan timur": "samarinda",
		"kalimantan tengah": "palangkaraya",
		"kalimantan barat":  "pontianak"}

	for key, value := range kamus {
		fmt.Println("Key: ", key, " Value: ", value)
	}
}
