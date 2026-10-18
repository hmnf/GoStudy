package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type UserStore struct {
	mu     sync.Mutex
	users  []User
	nextID int
}
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
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
	userStore := UserStore{
		mu:     sync.Mutex{},
		users:  users,
		nextID: users[len(users)-1].ID + 1,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", userStore.getUsers)
	mux.HandleFunc("POST /users", userStore.addUser)
	http.ListenAndServe(":8080", mux)
}

func (s *UserStore) getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(&s.users)
}

func (s *UserStore) addUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	s.users = append(s.users, User{
		ID:   s.nextID,
		Name: user.Name,
		Age:  user.Age,
	})
	s.nextID++
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
