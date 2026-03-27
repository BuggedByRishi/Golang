package main

import "fmt"

type Speaker interface { // This is the Interface
	Speak()
}

type Dog struct{} // Thsi is the structure

func (d Dog) Speak() {
	fmt.Println("Woof!")
}

type cat struct{}

func (c cat) Speak() {
	fmt.Println("Meow!")
}

func main() {
	var s Speaker

	s = Dog{}
	s.Speak()

	s = cat{}
	s.Speak()
}
