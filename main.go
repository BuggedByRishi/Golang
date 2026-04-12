package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 1; i <= 3; i++ {
		fmt.Println("Hello", i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go sayHello() // runs in background

	for i := 1; i <= 3; i++ {
		fmt.Println("Main", i)
		time.Sleep(time.Millisecond * 500)
	}
}
