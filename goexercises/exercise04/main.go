package main

import "fmt"

func main() {
	/*
		Problem

		Take an integer as input

		Print:
		"Even number" OR
		"Odd number"

		Constraint:
		Use only if / else
	*/
	var number int
	fmt.Println("Enter a number:")
	fmt.Scan(&number)
	if(number%2==0) {
		fmt.Printf("%d is an Even number", number)
	} else {
		fmt.Printf("%d is an Odd number", number)
	}
}