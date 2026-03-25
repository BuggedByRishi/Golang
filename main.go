package main

import (
	"fmt"
	"time"
)

func Message() {
	fmt.Print("If you are reading this it's too late")
}

func main() {
	go Message()

	time.Sleep(5 * time.Second)
}
