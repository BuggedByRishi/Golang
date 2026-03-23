package main

import "fmt"

func main() {
	power := 1000 // The variable should always be used in the code somewhere
	value := 100
	fmt.Printf("The value is: %d\n", value)
	fmt.Printf("default power is %d\n", power)
	name, power := "Rishi", 99999
	fmt.Printf("%s's power is over %d\n", name, power)
}
