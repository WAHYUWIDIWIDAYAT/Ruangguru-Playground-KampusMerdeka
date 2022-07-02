package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// sign dan welcome menggunakan JWT kedalam cookie
func main() {
	// start the server on port 8000
	fmt.Println("Starting Server at port :8000")
	log.Fatal(http.ListenAndServe(":8000", Routes()))
}

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		var creds Credentials
		// Task: JSON body diconvert menjadi creditial struct & return bad request ketika terjadi kesalahan decoding json:

		//beginanswer
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			// return bad request ketika terjadi kesalahan decoding json
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		//endanswer

		// Task: Ambil password dari username yang dipakai untuk login & return unauthorized jika password salah

		//beginanswer
		expectedPassword, ok := users[creds.Username]
		// return unauthorized jika password salah
		if !ok || expectedPassword != creds.Password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		//endanswer

		//Task: 1. Deklarasi expiry time untuk token jwt
		// 		2. Buat claim menggunakan variable yang sudah didefinisikan diatas
		//      3. Expiry time menggunakan time millisecond

		//beginanswer
		// Deklarasi expiry time untuk token jwt
		expirationTime := time.Now().Add(5 * time.Minute)
		// Buat claim menggunakan variable yang sudah didefinisikan diatas
		claims := &Claims{
			Username: creds.Username,
			StandardClaims: jwt.StandardClaims{
				// expiry time menggunakan time millisecond
				ExpiresAt: expirationTime.Unix(),
			},
		}
		//endanswer

		//Task: 1. Buat token menggunakan encoded claim dengan salah satu algoritma yang dipakai
		// 		2. Buat jwt string dari token yang sudah dibuat menggunakan JWT key yang telah dideklarasikan
		//      3. return internal error ketika ada kesalahan ketika pembuatan JWT string

		//beginanswer
		// Buat token menggunakan encoded claim dengan salah satu algoritma yang dipakai
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		// Buat jwt string dari token yang sudah dibuat menggunakan JWT key yang telah dideklarasikan
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			// return internal error ketika ada kesalahan ketika pembuatan JWT string
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		//endanswer

		//Task: Set token string kedalam cookie response

		//beginanswer
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
		//endanswer
	})

	mux.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		// Task: 1. Ambil token dari cookie yang dikirim ketika request
		//		 2. Buat return unauthorized ketika token kosong
		//		 3. Buat return bad request ketika field token tidak ada

		//beginanswer
		// Ambil token dari cookie yang dikirim ketika request
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				// return unauthorized ketika token kosong
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// return bad request ketika field token tidak ada
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		//endanswer

		// Task: Ambil value dari cookie token

		//beginanswer
		tknStr := c.Value

		// xxx.yyy.zzz

		//endanswer

		// Task: Deklarasi variable claim

		//beginanswer
		claims := &Claims{}
		//endanswer

		//Task: parse JWT token ke dalam claim

		//beginanswer
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		//endanswer

		//Task: return unauthorized ketika token sudah tidak valid
		// (biasanya karna token expired)

		//beginanswer
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		//endanswer

		// Task: return data dalam claim, seperti username yang telah didefinisikan

		//beginanswer
		w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
		//endanswer
	})

	return mux
}
