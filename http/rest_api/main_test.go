package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"slices"
	"sync"
	"testing"
)

func TestGetUsers(t *testing.T) {
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
		nextID: 4,
	}
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()

	userStore.getUsers(rec, req)
	resp := rec.Result()

	if resp.StatusCode != 200 {
		t.Fatalf("wrong status code %v, want %v", resp.StatusCode, http.StatusOK)
	}

	if resp.Header.Get("Content-Type") != "application/json" {
		t.Fatalf("wrong content type %v, want %v", resp.Header.Get("Content-Type"), "application/json")
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	var usersRes []User
	err = json.Unmarshal(data, &usersRes)
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(userStore.users, usersRes) {
		t.Fatalf("get data %v != want data %v", usersRes, userStore.users)
	}
	resp.Body.Close()
}
