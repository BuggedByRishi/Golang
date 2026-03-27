package main

import (
	"errors"
	"fmt"
)

func withdraw(balance, amount int) (int, error) {
	if amount > balance {
		return balance, errors.New("Insufficient Balancde")
	}
	return balance - amount, nil
}

func main() {
	amount := 10000000000000
	balance := 1000000000

	newBalance, err := withdraw(balance, amount)

	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Current Balance:", newBalance)
		return
	}

	fmt.Println("Withdrawal successful!")
	fmt.Println("Remaining Balance:", newBalance)
}
