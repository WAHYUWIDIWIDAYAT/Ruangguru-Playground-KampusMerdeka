package main

import (
	"encoding/json"
)

// TODO: answer here
type Ruang struct {
	RuangTamu Items `json:"ruangTamu"`
}

type Item struct {
	Nama   string `json:"nama"`
	Jumlah int    `json:"jumlah"`
	Warna  string `json:"warna"`
	Ukuran Ukuran `json:"ukuran"`
}

type Items struct {
	Items []Item `json:"items"`
}

type Ukuran struct {
	Panjang string `json:"panjang"`
	Tinggi  string `json:"tinggi"`
}

func (r Ruang) EncodeJSON() string {
	// TODO: answer here
	// encode
	data, _ := json.Marshal(r)
	// convert to string
	return string(data)

}

func NewRuang(r Ruang) Ruang {
	return r
}

func main() {

	// masukan data ke struct Item
	mejaMeja := []Item{

		{
			Nama:   "Meja",
			Jumlah: 20,
			Warna:  "Coklat",
			Ukuran: Ukuran{
				Panjang: "50 cm",
				Tinggi:  "25 cm",
			},
		},
		{
			Nama:   "Kursi",
			Jumlah: 1,
			Warna:  "Hitam",
			Ukuran: Ukuran{
				Panjang: "70 cm",
				Tinggi:  "30 cm",
			},
		},
	}

	// masukan mejaMeja ke struct Items
	items := Items{
		mejaMeja,
	}

	ruang := Ruang{
		RuangTamu: items,
	}

	// encode ke funtion EncodeJSON
	encode := ruang.EncodeJSON()

	// print
	println(encode)
}
