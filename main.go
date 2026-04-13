package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 42 // Send value
	}()

	value := <-ch // Receive value
	fmt.Println("Received:", value)
}
