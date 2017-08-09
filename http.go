package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", earthHandler)
	http.HandleFunc("/mars", marsHandler)

	http.ListenAndServe(":8080", nil)
}

func earthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Earth!\n")
}

func marsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Mars!\n")
}
