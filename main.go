package main

import "fmt"

type Animal struct {
	name string
}

type Dog struct {
	Animal
	Breed string
}

func main() {
	p := Dog{
		Animal: Animal{name: "Sheru"},
		Breed:  "Street Dog",
	}

	fmt.Println(p.name, p.Breed)
}
