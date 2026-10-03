package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", GetUsers)
	mux.HandleFunc("POST /users", PostUsers)
	mux.HandleFunc("PUT /users", PutUsers)
	mux.HandleFunc("DELETE /users", DeleteUsers)
	http.ListenAndServe(":8080", mux)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method)
	fmt.Println("Path: ", r.URL.Path)
	io.WriteString(w, "GET /users -> Getting users")
}

func PostUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method)
	fmt.Println("Path: ", r.URL.Path)
	io.WriteString(w, "POST /users -> Creating users")
}

func PutUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method)
	fmt.Println("Path: ", r.URL.Path)
	io.WriteString(w, "PUT /users -> Updating users")
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method)
	fmt.Println("Path: ", r.URL.Path)
	io.WriteString(w, "DELETE /users -> Deleting users")
}
