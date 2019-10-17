package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>HELLO WORLD</h1>"))
	})
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>This is homepage</h1>"))
	})
	http.HandleFunc("/output", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			input := r.FormValue("input")
			w.Write([]byte(input))
			return
		}
		http.Error(w, "Bad request", http.StatusBadRequest)
	})
	fmt.Println("server running in port 8000")
	http.ListenAndServe(":8000", nil)
}