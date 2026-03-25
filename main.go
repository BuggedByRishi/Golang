package main

import "fmt"

func main() {
	x := 10
	p := &x

	*p = 20 // Modifying value

	fmt.Println(x)
}
