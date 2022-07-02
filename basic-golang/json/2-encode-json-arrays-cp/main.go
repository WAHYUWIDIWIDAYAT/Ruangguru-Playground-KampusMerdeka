package main

import "encoding/json"

// Buat string JSON dengan hasil seperti berikut
// [{"jenis":"Meja Lipat","warna":"Coklat","jumlah":40,"deskripsi":"meja untuk belajar"},{"jenis":"Meja Hijau","warna":"Hijau","jumlah":10,"deskripsi":"meja untuk pengadilan"}]

type Meja struct {
	// TODO: answer here
	Jenis     string `json:"jenis"`
	Warna     string `json:"warna"`
	Jumlah    int    `json:"jumlah"`
	Deskripsi string `json:"deskripsi"`
}

type Items struct {
	MejaMeja []Meja
}

func (m Items) EncodeJSON() string {
	// TODO: answer here
	// encodeJSON
	data, _ := json.Marshal(m.MejaMeja)
	return string(data)

}

func NewMeja(m Items) Items {
	return m
}

func main() {
	// isi data ke struct
	items := Items{
		MejaMeja: []Meja{
			{
				Jenis:     "Meja Lipat",
				Warna:     "Coklat",
				Jumlah:    40,
				Deskripsi: "meja untuk belajar",
			},
			{
				Jenis:     "Meja Hijau",
				Warna:     "Hijau",
				Jumlah:    10,
				Deskripsi: "meja untuk pengadilan",
			},
		},
	}

	// masukan data ke fungsi NewMeja
	meja := NewMeja(items)

	// masukan nilai ke fungsi EncodeJSON
	result := meja.EncodeJSON()

	// cetak hasil
	println(result)

}
