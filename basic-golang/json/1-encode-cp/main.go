package main

import "encoding/json"

// Dari contoh yang telah diberikan, cobalah untuk melakukan encode struct menjadi json.
// Lengkapi function EncodeToJson agar dapat mengembalikan nilai byte hasil dari encode objek Leaderboard.
// Modifikasi struct UserRank sehingga field Name menjadi name ,field Rank menjadi rank, dan field Email tidak ikut untuk diencode.

type UserRank struct {
	Name  string `json:"name"`
	Email string `json:"-"`
	Rank  int    `json:"rank"`
}

type Leaderboard struct {
	Users []*UserRank
}

func EncodeToJson(leaderboard Leaderboard) ([]byte, error) {
	// TODO: answer here
	data, err := json.Marshal(leaderboard)
	return data, err
}

func main() {

	// masukan data ke struct
	users := []*UserRank{
		{
			Name:  "Roger",
			Email: "roger@ruangguru.com",
			Rank:  1,
		},
		{
			Name:  "Tony",
			Email: "tony@ruangguru.com",
			Rank:  2,
		},
		{
			Name:  "Bruce",
			Email: "bruce@ruangguru.com",
			Rank:  3,
		},
		{
			Name:  "Natasha",
			Email: "natasha@ruangguru.com",
			Rank:  4,
		},
		{
			Name:  "Clint",
			Email: "clint@ruangguru.com",
			Rank:  5,
		},
	}

	// isi data struct ke leaderboard
	leaderboard := Leaderboard{
		Users: users,
	}

	// encode data leaderboard ke json
	data, err := EncodeToJson(leaderboard)

	// cek error
	if err != nil {
		panic(err.Error())
	}

	// tampilkan data json
	println(string(data))

}
