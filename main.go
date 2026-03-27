package main

import (
	"fmt"
	"net/http"
)

func divideHandler(w http.ResponseWriter, r *http.Request) {
	a := 10
	b := 0

	if b == 0 {
		http.Error(w, "Cannot divide by Zero", http.StatusBadRequest)
		return
	}

	result := a / b
	fmt.Fprintln(w, "Result:", result)
}

func main() {
	http.HandleFunc("/divide", divideHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
