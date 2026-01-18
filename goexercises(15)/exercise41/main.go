package main

import "fmt"

func main() {
	/*
		Problem:
			Take an integer N from the user

		Task:
			Create a function that returns the sum of numbers from 1 to N

		Print:
			"Total sum: <value>"

		Constraint:
			Use for loop inside function
	*/
	fmt.Println("Enter a number for natural number sum")
	var num int
	fmt.Scan(&num)

	fmt.Println("Total sum:", sumN(num))
}

func sumN(n int) int {
	// sum := 0
	// for i := 1; i <= n; i++ {
	// 	sum += i
	// }
	return n*(n+1)/2
}