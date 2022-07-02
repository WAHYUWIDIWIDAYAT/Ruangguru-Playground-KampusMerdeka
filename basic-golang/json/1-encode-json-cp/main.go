package main

import (
	"encoding/json"
	"fmt"
)

// Dari struct dan nama field yang sama dari contoh
// Buat string JSON dengan hasil seperti berikut
// {"jenis":"Meja Belajar","color":"green","jumlah":2}

type Meja struct {
	// TODO: answer here
	Jenis  string `json:"jenis"`
	Warna  string `json:"color"`
	Jumlah int    `json:"jumlah"`
}

func (m Meja) EncodeJSON() string {
	// TODO: answer here
	// konversti sturct ke json
	data, _ := json.Marshal(m)
	return string(data)

}

func NewMeja(m Meja) Meja {
	return m
}

func main() {
	// masukan data ke struct
	var meja Meja
	meja.Jenis = "Meja Belajar"
	meja.Warna = "green"
	meja.Jumlah = 2

	// isi data struct ke leaderboard
	result := meja.EncodeJSON()

	// cek hasil
	fmt.Println(result)
	// fmt.Println(meja.EncodeJSON())

}
