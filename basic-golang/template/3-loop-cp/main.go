package main

import (
	"bytes"
	"html/template"
)

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan looping pada template.
// Lengkapi function ExecuteToByteBuffer dan textTemplate sehingga template dapat mencetak objek Leaderboard dengan ketentuan:
// Lakukan looping untuk setiap user
// Pada setiap loop, cetak "Peringkat ke-n: Nama", contoh: "Peringkat ke-1: Roger"

type UserRank struct {
	Name  string
	Email string
	Rank  int
}

type Leaderboard struct {
	Users []*UserRank
}

func ExecuteToByteBuffer(leaderboard Leaderboard) ([]byte, error) {
	var textTemplate string
	// TODO: answer here
	Users := leaderboard.Users
	textTemplate = `Peringkat ke-{{.Rank}}: {{.Name}}`
	t := template.New("textTemplate")
	t, err := t.Parse(textTemplate)
	if err != nil {
		return nil, err
	}
	var b bytes.Buffer
	for _, user := range Users {
		err = t.Execute(&b, user)
		if err != nil {
			return nil, err
		}
	}
	return b.Bytes(), nil

}
