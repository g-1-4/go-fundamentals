package main

import "fmt"

func main() {

	/*
		Problem
		Create a menu-driven program.

		Menu
		1. Add number
		2. View numbers
		3. Exit


		Task
		Store numbers in a slice
		Use for loop and switch
	*/
	mySlice := []int{}

	for {
		fmt.Println("\nMenu")
		fmt.Println("1. Add number")
		fmt.Println("2. View numbers")
		fmt.Println("3. Exit")
		fmt.Println("Enter a number in menu!")

		var choice int
		fmt.Scan(&choice)
		
		switch choice {
		case 1:
			var num int
			fmt.Println("Enter the number you want to add")
			fmt.Scan(&num)
			mySlice = append(mySlice, num)

		case 2:
			fmt.Println(mySlice)

		case 3:
			fmt.Println("Exiting the program...")
			return

		default:
			fmt.Println("Enter a number from menu")
		}
	}

}