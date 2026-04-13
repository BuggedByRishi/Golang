package main

import (
	"fmt"
	"net/http"
)

type User struct {
	ID   int
	Name string
}

type UserRepository interface {
	GetUserByID(id int) (User, error)
}

type InMemoryUserRepo struct{}

func (r InMemoryUserRepo) GetUserByID(id int) (User, error) {
	return User{ID: 1, Name: "Hrushikesh Kakulte"}, nil
}

type UserService struct {
	repo UserRepository
}

func (s UserService) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	user, err := s.repo.GetUserByID(1)

	if err != nil {
		fmt.Fprintln(w, "User not found")
		return
	}

	fmt.Fprintln(w, "User:", user.Name)
}

func main() {
	repo := InMemoryUserRepo{}
	service := UserService{repo: repo}

	http.HandleFunc("/user", service.HandleGetUser)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
