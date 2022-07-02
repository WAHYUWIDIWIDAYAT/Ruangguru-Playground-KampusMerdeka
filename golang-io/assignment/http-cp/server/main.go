package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, world!"))
	})
	mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		data := r.URL.Query().Get("data")
		w.Write([]byte(data))
	})
	mux.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		a := r.URL.Query().Get("a")
		aInt, err := strconv.Atoi(a)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		b := r.URL.Query().Get("b")
		bInt, err := strconv.Atoi(b)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		sum := aInt + bInt
		w.Write([]byte(fmt.Sprint(sum)))
	})
	mux.HandleFunc("/hellojson", func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Hello, world!"}`))
	})

	return mux
}
func main() {
	http.ListenAndServe(":8080", Routes())
}
