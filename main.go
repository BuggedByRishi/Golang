package main

import (
	"errors"
	"fmt"
)

func checkEven(n int) error {
	if n%2 != 0 {
		return errors.New("not an even number")
	}
	return nil
}

func main() {
	num := 7

	err := checkEven(num)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Number is even: ", num)
}
