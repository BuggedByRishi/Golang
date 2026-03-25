package main

import "fmt"

type Person struct {
	Name string
}

func changeName(p *Person) {
	p.Name = "Hrushikesh"
}

func main() {
	p := Person{Name: "Rishi"}
	changeName(&p)

	fmt.Print(p.Name)
}
