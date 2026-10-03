package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", h1)
	mux.HandleFunc("GET /hello?name=Ars", h2)
	mux.HandleFunc("GET /hello?name=Alex", h3)
	http.ListenAndServe(":8080", mux)
}

func h1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	fmt.Println("path: ", r.URL)
	io.WriteString(w, "Hello, stranger!")
}

func h2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	fmt.Println("path: ", r.URL)
	fmt.Println("path: ", r.URL.Query().Get("name"))
	io.WriteString(w, "Hello, Arseniy!")
}

func h3(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	fmt.Println("path: ", r.URL)
	fmt.Println("path: ", r.URL.Query().Get("name"))
	io.WriteString(w, "Hello, Alex!")
}
