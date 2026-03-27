package main

import "fmt"

type Payment interface {
	Pay(amount int)
}

type upi struct{}

func (u upi) Pay(amount int) {
	fmt.Println("Paid via UPI: ", amount)
}

type card struct{}

func (c card) Pay(amount int) {
	fmt.Println("Paid via Card: ", amount)
}

func processPayment(p Payment) {
	p.Pay(10000000000)
}

func main() {
	u := upi{} // Object creation			// This is for UPI payment
	processPayment(u)

	c := card{}       // Object creation
	processPayment(c) // This if for CARD PAyment
}
