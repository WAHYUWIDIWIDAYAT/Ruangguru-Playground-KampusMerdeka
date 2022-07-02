package main

import (
	"net/http"
)

var names = []string{
	"Tony",
	"Roger",
	"Bruce",
	"Stephen",
}

func IsNameExists(name string) bool {
	for _, n := range names {
		if n == name {
			return true
		}
	}

	return false
}

func GetNameHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		name := r.URL.Query().Get("name")
		if !IsNameExists(name) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	// TODO: answer here
	mux.HandleFunc("/name", GetNameHandler())
	return mux
}