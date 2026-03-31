package main

import "fmt"

func sendData(ch chan string) {
	ch <- "Hello from goroutine"
}

func main() {
	ch := make(chan string)

	go sendData(ch)

	msg := <-ch // receive data
	fmt.Println(msg)
}
