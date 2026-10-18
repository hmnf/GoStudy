package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /user", handler)
	http.ListenAndServe(":8080", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{
			Id:   1,
			Name: "Ars",
			Age:  18,
		},
		{
			Id:   2,
			Name: "Oleg",
			Age:  30,
		},
		{
			Id:   3,
			Name: "Olya",
			Age:  40,
		},
	}
	json.NewEncoder(w).Encode(users)
}
