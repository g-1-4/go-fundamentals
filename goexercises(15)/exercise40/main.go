package main

import "fmt"

func main() {
	/*
		Problem:
			Take a string from the user

		Task:
			Create a function that returns the length of the string

		Print:
			"Length is: <number>"

		Constraint:
			No built-in len() inside main
	*/
	fmt.Println("Enter a string to know its length")
	var str string
	fmt.Scan(&str)
	fmt.Println("length is:", lengthString(str))
}

func lengthString(s string) int {
	count := 0
	for range s {
		count += 1
	}
	return count
}