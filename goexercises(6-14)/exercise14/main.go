package main

import "fmt"

func main() {
	/*
		Problem
		Create an empty slice of integers.

		Task
		Add 5 numbers using append

		Print:
		Slice elements
		Length
		Capacity

		Constraint
		No predefined size
	*/
	mySlice := make([]int, 0)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 3)
	mySlice = append(mySlice, 4)
	mySlice = append(mySlice, 5)

	for _, i := range mySlice {
		fmt.Println(i)
	}

	fmt.Printf("The length of the slice is %d\n", len(mySlice))
	fmt.Printf("The capacity of slice is %d\n", cap(mySlice))
}