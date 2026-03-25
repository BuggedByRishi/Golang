package main

import "fmt"

func main() {
	x := 10
	p := &x

	fmt.Print("The value of x is : ", x)          // The actual value
	fmt.Print("\nThe address of x is : ", p)      // Address of the value
	fmt.Print("\nValues using pointer is : ", *p) // The value Pointe holds
}
