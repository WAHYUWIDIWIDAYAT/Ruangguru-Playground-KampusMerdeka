package main

import (
	"fmt"
)

type Barang struct {
	Nama string `json:"nama"`
	Harga int `json:"harga"`
	JenisBarang string `json:"jenis_barang"`
}

func main(){
	json := `{"nama":"Buku","harga":10000,"jenis_barang":"Buku"}`
	jsonData := []byte(jsonString)
	var barang Barang

	err := json.Unmarshal(jsonData, &barang)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(barang.Nama)
	fmt.Println(barang.Harga)
	fmt.Println(barang.JenisBarang)
	
}