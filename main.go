package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the GO server")
	fmt.Fprintln(w, "IF YOU ARE READING THIS IT'S TOO LATE")
}

func main() {
	http.HandleFunc("/", homeHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
