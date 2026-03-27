package main

import "fmt"

type shape interface { // 2
	Area() int // This is the interface
}

type rectangle struct { // 3
	width, height int // This is the structure
}

func (r rectangle) Area() int { // 4
	return r.height * r.width
}

func printArea(s shape) { // 5
	fmt.Println(s.Area())
}

func main() { // 1
	r := rectangle{10, 20}
	printArea(r)
}
