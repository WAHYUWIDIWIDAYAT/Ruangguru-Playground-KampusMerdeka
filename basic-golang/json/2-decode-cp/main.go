package main

import "encoding/json"

// Dari contoh yang telah diberikan, cobalah untuk melakukan decode dari json menjadi objek dari struct.
// Mengambil kasus pada encode, lengkapi function DecodeToLeaderboard agar json dapat di decode menjadi objek Leaderboard

type UserRank struct {
	Name  string
	Email string `json:"-"`
	Rank  int
}

type Leaderboard struct {
	Users []*UserRank
}

func DecodeToLeaderboard(jsonData []byte) (Leaderboard, error) {
	// TODO: answer here
	// inisiasi leaderboard
	var leaderboard Leaderboard

	// decode json ke struct
	err := json.Unmarshal(jsonData, &leaderboard)

	// return
	return leaderboard, err
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

	// jadiin user jason
	jsonData, _ := json.Marshal(users)

	// masukan json ke function DecodeToLeaderboard
	DecodeToLeaderboard(jsonData)

}
