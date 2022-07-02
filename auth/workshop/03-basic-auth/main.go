package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// start the server on port 8000
	fmt.Println("Starting Server at port :8000")
	log.Fatal(http.ListenAndServe(":8000", Routes()))
}

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	// encrypt token dari username menggunakan bcrypt kemudian kirim ke user kedalam cookie
	mux.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		cookieName := "token"
		var creds Credentials

		// Task:  Buat JSON body diconvert menjadi credential struct & return bad request ketika terjadi kesalahan decoding json

		//beginanswer
		// JSON body diconvert menjadi credential struct
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			// return bad request ketika terjadi kesalahan decoding json
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		//endanswer

		// Task: Ambil password dari username yang dipakai untuk login

		//beginanswer
		// Ambil password dari username yang dipakai untuk login
		expectedPassword, ok := users[creds.Username]
		//endanswer

		// Task: return unauthorized jika password salah

		//beginanswer
		// return unauthorized jika password salah
		if !ok || expectedPassword != creds.Password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		//endanswer

		// Task: 1. Buat token string menggunakan bcrypt dari credential username
		// 		 2. return internal server error ketika terjadi kesalahan encrypting token

		//beginanswer
		// Buat token string menggunakan bcrypt dari credential username
		tokenString, err := bcrypt.GenerateFromPassword([]byte(creds.Username), bcrypt.DefaultCost)
		if err != nil {
			// return internal server error ketika terjadi kesalahan encrypting token
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		//endanswer

		// Task: Set token string kedalam cookie response

		//beginanswer
		// Set token string kedalam cookie response
		expirationTime := time.Now().Add(5 * time.Minute)
		http.SetCookie(w, &http.Cookie{
			Name:    cookieName,
			Value:   string(tokenString),
			Expires: expirationTime,
		})
		//endanswer
	})

	return mux
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
