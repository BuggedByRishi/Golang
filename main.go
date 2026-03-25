package main

import "fmt"

func update(num *int) {
	*num = 50
}

func main() {
	x := 10

	update(&x)

	fmt.Print(x)
}
