package main

import "fmt"

func main() {
	/*
		Problem:
			Take a number from the user

		Task:
			Create a function that returns factorial of the number

		Print:
			"Factorial is: <value>"

		Constraint:
			No recursion
	*/
	fmt.Println("Enter a number for natural number sum")
	var num int
	fmt.Scan(&num)

	if res := fact(num); res == -1 {
		fmt.Println("Factorial not defined for negative numbers")
	} else{
		fmt.Println("Factorial is:", res)
	}
}

func fact(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 1
	}
	product := 1
	for i := n; i > 0; i-- {
		product *= i
	} 
	return  product
}