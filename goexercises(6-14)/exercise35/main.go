package main

import "fmt"

func main() {
	/*
		Problem
		Given a slice of integers.

		Task
		Use pointer to slice
		Multiply all elements by 2
		Print updated slice
	*/

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	ptr := &nums

	for i := range *ptr {
		(*ptr)[i] = (*ptr)[i] * 2
	}

	fmt.Println("The updated slice is:", *ptr)
}