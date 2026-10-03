package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", handler)
	http.ListenAndServe(":8080", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	fmt.Println("error: ", err)

	fmt.Println("user: ", user.Name, "age ", user.Age, "created")
}
