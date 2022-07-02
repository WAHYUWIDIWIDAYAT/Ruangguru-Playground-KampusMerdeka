package main

import (
	"bytes"
	"fmt"
	"html/template"
)

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan condition pada template.
// Lengkapi function ExecuteToByteBuffer dan variabel textTemplate sehingga memiliki keluaran:
// 1. Jika saldo sama dengan 0, cetak "Akun Tony dengan Nomor 1002321 tidak memiliki saldo."
// 2. Jika saldo lebih dari 0, cetak "Akun Tony dengan Nomor 1002321 memiliki saldo sebesar $1000."

type Account struct {
	Name    string
	Number  string
	Balance int
}

func ExecuteToByteBuffer(account Account) ([]byte, error) {
	var textTemplate string
	textTemplate = `Akun {{.Name}} dengan Nomor {{.Number}} {{if .Balance}}memiliki saldo sebesar ${{.Balance}}{{else}}tidak memiliki saldo{{end}}.`
	t := template.New("textTemplate")
	t, err := t.Parse(textTemplate)
	if err != nil {
		return nil, err
	}
	var b bytes.Buffer
	err = t.Execute(&b, account)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func main() {
	var account Account
	account.Name = "Tony"
	account.Number = "1002321"
	account.Balance = 1000
	b, err := ExecuteToByteBuffer(account)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

}
