package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Numbers: ", i)
		time.Sleep(time.Millisecond * 300)
	}
}

func printLetters() {
	for i := 'A'; i <= 'Z'; i++ {
		fmt.Println("Letters: ", string(i))
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	go printLetters()
	go printNumbers()

	time.Sleep(2 * time.Second)
}
