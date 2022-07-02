package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Authorization menggunakan claim JWT
func main() {
	// start the server on port 8000
	fmt.Println("Starting Server at port :8000")
	log.Fatal(http.ListenAndServe(":8000", Routes()))
}

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		var creds Credentials
		// JSON body diconvert menjadi creditial struct & return bad request ketika terjadi kesalahan decoding json

		//beginanswer
		// JSON body diconvert menjadi creditial struct
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			// return bad request ketika terjadi kesalahan decoding json
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		//endanswer

		// Task: Ambil password dari username yang dipakai untuk login & return unauthorized jika password salah

		//beginanswer
		// Ambil password dari username yang dipakai untuk login
		expectedPassword, ok := users[creds.Username]
		// return unauthorized jika password salah
		if !ok || expectedPassword.Password != creds.Password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		//endanswer

		// Task: 1. Deklarasi expiry time untuk token jwt
		// 		 2. Buat claim menggunakan variable yang sudah didefinisikan diatas
		//       3. Expiry time menggunakan time millisecond

		//beginanswer
		// Deklarasi expiry time untuk token jwt
		expirationTime := time.Now().Add(5 * time.Minute)
		// Buat claim menggunakan variable yang sudah didefinisikan diatas
		claims := &Claims{
			Username: creds.Username,
			Role:     expectedPassword.Role,
			StandardClaims: jwt.StandardClaims{
				// expiry time menggunakan time millisecond
				ExpiresAt: expirationTime.Unix(),
			},
		}
		//endanswer

		// Task: Buat token menggunakan encoded claim dengan salah satu algoritma yang dipakai

		//beginanswer
		// Buat token menggunakan encoded claim dengan salah satu algoritma yang dipakai
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		//endanswer

		// Task: 1. Buat jwt string dari token yang sudah dibuat menggunakan JWT key yang telah dideklarasikan
		//       2. Buat return internal error ketika ada kesalahan ketika pembuatan JWT string

		//beginanswer
		// Buat jwt string dari token yang sudah dibuat menggunakan JWT key yang telah dideklarasikan
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			// return internal error ketika ada kesalahan ketika pembuatan JWT string
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		//endanswer

		// Task: Set token string kedalam cookie response

		//beginanswer
		// Set token string kedalam cookie response
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
		//endanswer
	})

	mux.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		// Task: 1. Ambil token dari cookie yang dikirim ketika request
		//       2. return unauthorized ketika token kosong
		//       3. return bad request ketika field token tidak ada

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
		// Ambil value dari cookie token
		tknStr := c.Value
		//endanswer

		// Task: Deklarasi variable claim

		//beginanswer
		// Deklarasi variable claim
		claims := &Claims{}
		//endanswer

		// Task: 1. Parse JWT token ke dalam claim
		//       2. return unauthorized ketika ada kesalahan ketika parsing token
		//	     3. return bad request ketika field token tidak ada

		//beginanswer
		//	parse JWT token ke dalam claim
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				// return unauthorized ketika ada kesalahan ketika parsing token
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// return bad request ketika field token tidak ada
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		//endanswer

		// Task: return unauthorized ketika token sudah tidak valid (biasanya karna token expired)

		//beginanswer
		//return unauthorized ketika token sudah tidak valid (biasanya karna token expired)
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		//endanswer

		// Task: return unauthorized ketika role user tidak sesuai dengan role admin

		//beginanswer
		//return unauthorized ketika role user tidak sesuai dengan role admin
		if claims.Role != "admin" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		//endanswer

		// Task: return data dalam claim, seperti username yang telah didefinisikan

		//beginanswer
		// return data dalam claim, seperti username yang telah didefinisikan
		w.Write([]byte(fmt.Sprintf("Welcome Admin %s!", claims.Username)))
		//endanswer
	})

	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		// Task: 1. Ambil token dari cookie yang dikirim ketika request
		//       2. return unauthorized ketika token kosong
		//       3. return bad request ketika field token tidak ada

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
		// Ambil value dari cookie token
		tknStr := c.Value
		//endanswer

		// Task: Deklarasi variable claim

		//beginanswer
		// Deklarasi variable claim
		claims := &Claims{}
		//endanswer

		// Task: 1. Parse JWT token ke dalam claim
		//       2. return unauthorized ketika ada kesalahan ketika parsing token
		//	     3. return bad request ketika field token tidak ada

		//beginanswer
		//parse JWT token ke dalam claim
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				// return unauthorized ketika ada kesalahan ketika parsing token
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// return bad request ketika field token tidak ada
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		//endanswer

		// Task: return unauthorized ketika token sudah tidak valid (biasanya karna token expired)

		//beginanswer
		//return unauthorized ketika token sudah tidak valid (biasanya karna token expired)
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		//endanswer

		// Task: return data dalam claim, seperti username yang telah didefinisikan

		//beginanswer
		// return data dalam claim, seperti username yang telah didefinisikan
		w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
		//endanswer
	})

	return mux
}
