package main

import (
	"fmt"
	"net/http"
	"time"
)

// Dari contoh yang telah diberikan, buatlah http.ServeMux yang memiliki dua route
// Route pertama yaitu "/time" yang menghandle TimeHandler
// Route kedua yaitu "/hello" yang menghandle SayHelloHandler

func TimeHandler() http.HandlerFunc {
	t := time.Now()
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("%v, %v %v %v", t.Weekday(), t.Day(), t.Month(), t.Year())))
	}
	return handler
	// TODO: replace this
}

func SayHelloHandler() http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if name != "" {
			w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
			return
		}

		w.Write([]byte("Hello there"))
	}
	return handler
	// TODO: replace this
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/time", TimeHandler())
	mux.Handle("/hello", SayHelloHandler())
	// TODO: answer here
	return mux

}
