package main

import "fmt"

func main() {
	x := 10
	p := &x
	pp := &p

	fmt.Print(*p)  // returns memory of x
	fmt.Print(*pp) // returns memory of x
}
