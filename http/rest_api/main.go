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
	mux.HandleFunc("GET /users", handlerGet)
	mux.HandleFunc("POST /users", handlerPost)
	http.ListenAndServe(":8080", mux)
}

func handlerGet(w http.ResponseWriter, r *http.Request) {
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

func handlerPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	user.ID = 4
	http.StatusText(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
