package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Dari contoh yang telah diberikan, buatlah route untuk TimeHandler dan SayHelloHandler.
// Buatlah route "/time" pada TimeHandler dan "/hello" untuk SayHelloHandler dengan menggunakan ServeMux

var TimeHandler = func(writer http.ResponseWriter, request *http.Request) {
	t := time.Now()
	fmt.Fprintf(writer, "%v, %v %v %v", t.Weekday(), t.Day(), t.Month(), t.Year())

	// TODO: answer here
}

var SayHelloHandler = func(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	if name != "" {
		fmt.Fprintf(w, "Hello, %s!", name)
		return
	}
	if name == "" {
		fmt.Fprintf(w, "Hello there")
		return
	}

	// TODO: answer here
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/time", TimeHandler)
	mux.HandleFunc("/hello", SayHelloHandler)
	// TODO: answer here

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
