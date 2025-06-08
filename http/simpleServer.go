package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "Path: %s\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/home", LoggingMiddleware(http.HandlerFunc(handler)))
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		start := time.Now()
		next.ServeHTTP(w, r)
		end := time.Since(start).Seconds()
		fmt.Printf("%v\n", end)
	})
}
