package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// ================= MODEL =================
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ============== IN-MEMORY DB ==============
var users = []User{
	{ID: 1, Name: "Hrushikesh"},
	{ID: 2, Name: "John"},
}

// ================= HANDLERS =================

// GET /users
func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

// GET /users/{id}
func getUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(idStr)

	for _, user := range users {
		if user.ID == id {
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

// POST /users
func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	json.NewDecoder(r.Body).Decode(&newUser)

	newUser.ID = len(users) + 1
	users = append(users, newUser)

	json.NewEncoder(w).Encode(newUser)
}

// DELETE /users/{id}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(idStr)

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			fmt.Fprintln(w, "User deleted")
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

// ================= ROUTER =================
func router(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/users" && r.Method == http.MethodGet {
		getUsers(w, r)
		return
	}

	if r.URL.Path == "/users" && r.Method == http.MethodPost {
		createUser(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "/users/") && r.Method == http.MethodGet {
		getUser(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "/users/") && r.Method == http.MethodDelete {
		deleteUser(w, r)
		return
	}

	http.Error(w, "Route not found", http.StatusNotFound)
}

// ================= MAIN =================
func main() {
	http.HandleFunc("/", router)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
