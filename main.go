package main

import "fmt"

type Saiyan struct { // This is the structure defination (blueprint)
	Name  string
	Power int
}

func NewSaiyan(name string, power int) *Saiyan { // Pointer to the structure
	return &Saiyan{
		Name:  name,
		Power: power,
	}
}

func main() {
	goku := NewSaiyan("Goku", 9000)
	fmt.Println(goku.Name, goku.Power)
}
