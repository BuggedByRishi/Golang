package main

import "fmt"

func main() {
	done := make(chan bool)

	go func() {
		fmt.Println("Work Done")
		done <- true
	}()
	<-done
	fmt.Println("Main Finished")
}
