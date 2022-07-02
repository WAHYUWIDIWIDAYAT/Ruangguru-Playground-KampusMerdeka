package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// start the server on port 8000
	fmt.Println("Starting Server at port :8080")
	log.Fatal(http.ListenAndServe(":8080", Routes()))
}

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	// Membaca cookie token kemudian return cookie kedalam body response
	mux.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		cookieFieldName := "token"

		// Task:  1. Ambil token dari cookie yang dikirim ketika request
		// 		  2. return unauthorized ketika token kosong
		// 		  3. return bad request ketika field token tidak ada

		//beginanswer
		// Ambil token dari cookie yang dikirim ketika request
		c, err := r.Cookie(cookieFieldName)

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
		respond := fmt.Sprintf("Tokenmu adalah %s!", c.Value)
		responByte := []byte(respond)
		w.Write(responByte)
		//endanswer return w.Write([]byte(fmt.Sprintf("")))
	})

	return mux
}
