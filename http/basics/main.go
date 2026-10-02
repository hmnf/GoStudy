package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handler1)
	mux.HandleFunc("GET /hello", handler2)
	http.ListenAndServe(":8080", nil)
}

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Go HTTP server")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello!")
}
