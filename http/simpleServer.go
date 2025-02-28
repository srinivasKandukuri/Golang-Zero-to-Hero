package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "Path: %s\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
