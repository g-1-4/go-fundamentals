package main

import "fmt"

func main() {
	/*
		Problem:
			Create a function that takes an integer

		Task:
			Use defer to print the value AFTER function finishes

		Print:
			"Value is: <number>"

		Constraint:
			Defer must use function parameter
	*/
	myfunc(3)
}

func myfunc(n int) {
	defer fmt.Println("Value is:", n)
}