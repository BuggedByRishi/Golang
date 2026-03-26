package main

import "fmt"

type address struct {
	city string
}

type person struct {
	name    string
	address address
}

func main() {
	p := person{
		name: "Rishi",
		address: address{
			city: "Pune",
		},
	}
	fmt.Println(p.name, p.address.city)
}
