package main

import "fmt"

type Account struct {
	Name    string
	Balance int
}

func NewAccount(name string) *Account {
	return &Account{Name: name, Balance: 0}
}

func (a *Account) Deposit(amount int) {
	a.Balance += amount
}

func main() {
	acc := NewAccount("Hrushikesh")
	acc.Deposit(10000000000)

	fmt.Println(acc.Name, acc.Balance)
}
