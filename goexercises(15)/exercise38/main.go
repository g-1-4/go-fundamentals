package main

import "fmt"

func main() {
	/*
		Problem:
			Take two integers from the user

		Task:
			Create a function that returns their sum

		Print:
			"Sum is: <result>"

		Constraint:
			Function must return a value
			No global variables
	*/
	fmt.Println("Enter first number:")
	var num1 int
	fmt.Scan(&num1)
	fmt.Println("Enter second number:")
	var num2 int
	fmt.Scan(&num2)
	fmt.Println("Sum is:", add(num1, num2))
}

func add(a, b int) int {
	return a + b
}