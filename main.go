package main

import (
	"fmt"
	"time"
)

func display(str string) {
	for i := 0; i < 3; i++ {
		fmt.Println(str)
	}
}

func main() {
	go display("Hello, Goroutine!")

	time.Sleep(2 * time.Second)
	display("Hello, Main!")
}
