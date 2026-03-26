package main

import "fmt"

func main() {
	a := 10

	fmt.Println(a)
	outer()
}

func outer() {
	fmt.Println("This is the Outer function")
	inner()
}

func inner() {
	fmt.Println("This is the Inner function")
	innerMost()
}

func innerMost() {
	fmt.Println("This is the innerMost function")
}
