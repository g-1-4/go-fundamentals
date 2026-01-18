package main

import "fmt"

func main() {
	/*
		Problem
		Declare two integers.

		Task
		Swap values using pointers
		Print before and after

		Constraint
		No temporary variables
	*/
	num1 := 2
	num2 := 3
	ptra := &num1
	ptrb := &num2

	fmt.Printf("Before swap num1: %v, num2: %v\n", num1, num2)

	*ptra, *ptrb = *ptrb, *ptra

	fmt.Printf("after swap num1: %v, num2: %v\n", num1, num2)
}