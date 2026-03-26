package main

import "fmt"

func main() {
	p := struct {
		Name string
		Age  int
	}{
		Name: "John",
		Age:  25,
	}

	fmt.Println(p.Name, p.Age)
}
