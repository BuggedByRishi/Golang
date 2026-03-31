package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "If you are reading this it's too late..."
	}()

	print := <-ch

	fmt.Println(print)
}
