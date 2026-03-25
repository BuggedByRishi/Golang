package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}

	nums = append(nums, 4, 5, 6)

	for i, v := range nums { // Loop through Slice
		fmt.Println(i, v)
	}
}
