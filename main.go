package main

import "fmt"

func main() {
	x := 10
	p := &x

	fmt.Print("Value : ", x)
	fmt.Print("Address : ", p)
}
