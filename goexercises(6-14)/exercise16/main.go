package main

import "fmt"

func main() {
	/*
		Problem
		Given a slice of integers.

		Task
		Reverse the slice manually
		Print before and after

		Constraint
		No built-in reverse helpers
	*/
	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice before reversed", mySlice)

	for i := 0; i < len(mySlice)/2; i++ {
		j := len(mySlice)-i-1
		mySlice[i], mySlice[j] = mySlice[j], mySlice[i]
	}
	fmt.Println("Slice after reversed", mySlice)
}