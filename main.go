package main

import "fmt"

type Saiyan struct {
	Power int
}

func (s *Saiyan) PowerUp() {
	s.Power += 1000
}

func main() {
	goku := Saiyan{Power: 9000}
	goku.PowerUp()

	fmt.Println(goku.Power) // 10000
}
