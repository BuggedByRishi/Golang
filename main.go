package main

import (
	"fmt"
	"math"
)

type shape interface {
	Area() float64
}

type rectangle struct {
	width, height float64
}

func (r rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func calculateArea(s shape) float64 {
	return s.Area()
}

func main() {
	rect := rectangle{height: 5, width: 4}
	circle := Circle{radius: 2}

	fmt.Println(
		"Rectangle Area:",
		calculateArea(rect),
	)

	fmt.Println(
		"Circle Area:",
		calculateArea(circle),
	)

	mysterBox := interface{}(10)
	describeValue(mysterBox)
}

func describeValue(t interface{}) { // empty interface can accept any type
	fmt.Printf("Type: %T, Value: %v\n", t, t)
}
