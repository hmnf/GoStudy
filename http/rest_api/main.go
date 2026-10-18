package main

import (
	"encoding/json"
	"net/http"
	"strconv"
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
	mux.HandleFunc("GET /users/{id}", userStore.getUser)
	mux.HandleFunc("DELETE /users/{id}", userStore.deleteUser)
	http.ListenAndServe(":8080", mux)
}

func (s *UserStore) deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	for key, user := range s.users {
		if user.ID == id {
			s.users = append(s.users[:key], s.users[key+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Not found", http.StatusNotFound)
}

func (s *UserStore) getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, user := range s.users {
		if user.ID == id {
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.Error(w, "Not found", http.StatusNotFound)
}

func (s *UserStore) getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s.mu.Lock()
	defer s.mu.Unlock()
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
	s.mu.Lock()
	defer s.mu.Unlock()
	user.ID = s.nextID
	s.users = append(s.users, user)
	s.nextID++
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
