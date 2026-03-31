package main

func main() {
	ch := make(chan int, 30)

	ch <- 10
	ch <- 20
	ch <- 30

	println(<-ch)
	println(<-ch)
	println(<-ch)
}
