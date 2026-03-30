package main

import "fmt"

func main() {
	defer fmt.Println("World")
	fmt.Println("Hello")
}

/*  Now the order is reversed the output will still be the same because the defer statement will execute
	just before the ending curly braces of the main function, so the output will be:

	Output: Hello
        World
*/
