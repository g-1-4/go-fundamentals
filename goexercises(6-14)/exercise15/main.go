package main

import "fmt"

func main() {
	/*
		Problem
		Given a slice of integers.

		Task
		Create a new slice containing only even numbers
		Print both slices

		Constraint
		Use for and if
	*/

	givenSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resSlice := []int{}

	for _, i := range givenSlice {
		if i%2 == 0 {
			resSlice = append(resSlice, i)
		}
	}
	fmt.Println("Given slice is:", givenSlice)
	fmt.Println("Reslut Slice is:", resSlice)
}