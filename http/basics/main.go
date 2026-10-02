package main

import (
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handler1)
	mux.HandleFunc("GET /hello", handler2)
	http.ListenAndServe(":8080", mux)
}

func handler1(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "go http server")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello")
}
