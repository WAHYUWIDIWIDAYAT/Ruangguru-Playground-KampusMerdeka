package main

import (
	"encoding/json"
	"log"
)

// Dari contoh sebelumnya
// buat JSON string array nested seperti berikut
/*
[
{
		"jenis": "Meja Makan",
		"warna": "Coklat",
		"jumlah": 20,
		"ukuran": {
				"panjang": "50 cm",
				"tinggi": "25 cm"
		}
},
{
		"jenis": "Meja Lipat",
		"warna": "Hitam",
		"jumlah": 1,
		"ukuran": {
				"panjang": "70 cm",
				"tinggi": "30 cm"
		}
}
]
*/

type Ukuran struct {
	Panjang string `json:"panjang"`
	Tinggi  string `json:"tinggi"`
}

type Meja struct {
	Jenis  string `json:"jenis"`
	Warna  string `json:"warna"`
	Jumlah int    `json:"jumlah"`
	Ukuran Ukuran `json:"ukuran"`
}

type Items struct {
	MejaMeja []Meja
}

func (m Items) EncodeJSON() string {
	MejaToJSON, err := json.Marshal(m.MejaMeja)
	if err != nil {
		log.Println(err)
	}
	return string(MejaToJSON)
}

func NewMeja(m Items) Items {
	return m
}
