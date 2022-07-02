   
package main

import (
	"fmt"
	"net/http"
	"time"
)

// Dari contoh yang telah diberikan, buatlah route untuk TimeHandler dan SayHelloHandler.
// Buatlah route "/time" pada TimeHandler dan "/hello" untuk SayHelloHandler dengan menggunakan http.HandleFunc

var TimeHandler = func(writer http.ResponseWriter, request *http.Request) {
	// TODO: answer here
	t := time.Now()
	weekday := t.Weekday()
	day := t.Day()
	month := t.Month()
	year := t.Year()

	result := fmt.Sprintf("%v, %v %v %v", weekday, day, month, year)

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(result))
}

var SayHelloHandler = func(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	result := "Hello there"
	name := r.URL.Query().Get("name")
	if name != "" {
		result = fmt.Sprintf("Hello, %v!", name)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

func main() {
	http.HandleFunc("/time", TimeHandler)
	http.HandleFunc("/hello", SayHelloHandler)
	http.ListenAndServe("localhost:8080", nil)
}