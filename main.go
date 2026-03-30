package main

import "fmt"

func main() {
	fmt.Println("Hello")
	defer fmt.Println("World")
}

/* As function executes line by line this will also execute line by line
Output: Hello
        World
*/
