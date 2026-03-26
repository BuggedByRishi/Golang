package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Println("Hello,", p.Name)
}

func main() {
	p := Person{Name: "John"}
	p.Greet()
}
