package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p := person{name: "Rishi", age: 21}
	fmt.Println(p.name, p.age)
}
