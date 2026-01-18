package main

import "fmt"

func main() {
	/*
		Problem:
			Take an integer from the user

		Task:
			Create a function to check if the number is even or odd

		Print:
			"Even" OR "Odd"

		Constraint:
			Function must return bool
			if-else allowed
	*/
	fmt.Println("Enter a number to check if its even or odd")
	var num int
	fmt.Scan(&num)

	if evenOrOdd(num) {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

}

func evenOrOdd(a int) bool {
	if a%2 == 0 {
		return true
	} else {
		return false
	}
}