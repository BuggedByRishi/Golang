package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}

	nums = append(nums, 4, 5, 6)

	fmt.Print("\n", nums)    // After appending 4
	fmt.Print("\n", nums[0]) // Accessing element from array
	fmt.Print(len(nums))     // Length of array
	fmt.Print(cap(nums))     // Capacity or size of the array
}
