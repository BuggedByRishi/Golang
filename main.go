package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2) // Buffer channel(will execute 2 iterations of the loop before blocking)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()

	for i := 0; i <= 3; i++ {
		fmt.Println(<-ch)
	}
}
