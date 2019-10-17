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
	fmt.Println("server running in port 8000")
	http.ListenAndServe(":8000", nil)
}
