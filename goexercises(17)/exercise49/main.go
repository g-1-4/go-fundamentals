package main

import "fmt"

func main() {
	/*
		Problem:
			Print numbers from 1 to 3

		Task:
			Use defer inside loop

		Print:
			3
			2
			1

		Constraint:
			Use for loop
			No arrays or slices
	*/
	for i := 1; i <= 3; i++ {
		defer fmt.Println(i)
	}
}