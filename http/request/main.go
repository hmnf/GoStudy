package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", handler)
	http.ListenAndServe(":8080", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	fmt.Println(r.Method)
	fmt.Println(r.URL.Path)

	if name == "" {
		io.WriteString(w, "Hello, stranger!")
	} else {
		io.WriteString(w, fmt.Sprintf("Hello, %v!", name))
	}
}
