package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
}

/*
This code will follow the LIFO (Last In First Out)

Two <- One <- World

Two is the last element and will be the first to execute

*/
