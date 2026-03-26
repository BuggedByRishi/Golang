package main

import "fmt"

type calculation struct {
	num1, num2 int
}

func (c calculation) sum() int { // Method for sum
	return c.num1 + c.num2 // Methos is not like a function
}

func (c calculation) mul() int { // Method for multiplication
	return c.num1 * c.num2
}

func main() {
	cal := calculation{10, 20}
	fmt.Println("The sum of two numbers is : ", cal.sum())
	fmt.Println("The multiplication of two numbers is : ", cal.mul())
}
