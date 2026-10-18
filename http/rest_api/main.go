package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", handler)
	http.ListenAndServe(":8080", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := []User{
		{
			ID:   1,
			Name: "Ars",
			Age:  18,
		},
		{
			ID:   2,
			Name: "Oleg",
			Age:  30,
		},
		{
			ID:   3,
			Name: "Olya",
			Age:  40,
		},
	}
	json.NewEncoder(w).Encode(users)
}
