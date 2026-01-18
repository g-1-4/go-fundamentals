package main

import "fmt"

func main() {
	/*
		Problem
		Input a number (1–7).

		Task
		Print corresponding day name
		Print error for invalid input

		Constraint
		Use switch
	*/
	fmt.Println("Enter a number from 1 to 7")
	var num int
	fmt.Scan(&num)
	switch num {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid Input")
	}
}