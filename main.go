package main

import (
	"fmt"
	"time"
)

func Drake() {
	fmt.Println("If you are reading this it's too late")
}

func main() {
	go Drake()

	time.Sleep(1 * time.Second)
	fmt.Println("Inside the Main function ")
}
