package main

import (
	"bytes"
	"text/template"
)

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan function pada template.
// Lengkapi function CalculateScore yang berfungsi untuk menjumlahkan total score dari users
// Lengkapi function ExecuteToByteBuffer dan textTemplate sehingga template dapat mencetak objek Leaderboard dengan ketentuan:
// Lakukan looping untuk setiap user
// Pada setiap loop, cetak "Nama: Score", contoh: "Roger: 1000"
// Pada bagian terakhir, cetak "Total Score: total", contoh: "Total Score: 50000"

type UserRank struct {
	Name  string
	Email string
	Rank  int
	Score int
}

type Leaderboard struct {
	Users []*UserRank
}

func CalculateScore(leaderboard Leaderboard) int {
	// TODO: answer here
	// bikin variable total
	total := 0

	// looping
	for _, user := range leaderboard.Users {
		total += user.Score
	}
	return total
}

func ExecuteToByteBuffer(leaderboard Leaderboard) ([]byte, error) {

	// register function
	funcMap := template.FuncMap{
		"CalculateScore": CalculateScore,
	}

	// bikin kalimat yang akan ditampilkan
	textTemplate := `{{ range .Users }}{{ .Name }}: {{ .Score }}{{ end }}Total Score: {{ CalculateScore . }}`

	// bikin buffer
	var b bytes.Buffer

	// biking template
	tmpl, err := template.New("test").Funcs(funcMap).Parse(textTemplate)
	if err != nil {
		return nil, err
	}

	// generate template
	err = tmpl.Execute(&b, leaderboard)
	if err != nil {
		return nil, err
	}

	// kembalikan hasil
	return b.Bytes(), nil

}

func main() {
	// isi data ke dalam struct
	users := []*UserRank{
		{
			Name:  "Roger",
			Rank:  1,
			Score: 1000,
		},
		{
			Name:  "Tony",
			Rank:  2,
			Score: 900,
		},
		{
			Name:  "Bruce",
			Rank:  3,
			Score: 800,
		},
		{
			Name:  "Natasha",
			Rank:  4,
			Score: 700,
		},
		{
			Name:  "Clint",
			Rank:  5,
			Score: 600,
		},
	}

	// masukan users ke leaderboard
	leaderboardObject := Leaderboard{
		Users: users,
	}

	// generate output
	b, err := ExecuteToByteBuffer(leaderboardObject)
	if err != nil {
		panic(err)
	}

	// cetak output
	println(string(b))
}
