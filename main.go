package main

import (
	"fmt"
	"time"
)

func task() {
	fmt.Println("Task Started")
	time.Sleep(2 * time.Second)
	fmt.Println("Task finished")
}

func main() {
	go task()

	time.Sleep(3 * time.Second)
}
