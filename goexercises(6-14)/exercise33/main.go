package main

import "fmt"

func main() {
	/*
		Problem
		Declare an integer variable.

		Task
		Create pointer to it

		Print:
		Value
		Address
		Value via pointer
		Modify value using pointer
	*/

	num := 4

	numAddr := &num

	fmt.Println("The memory address of num is:", numAddr)
	fmt.Println("The value is:", *numAddr)

	*numAddr += 4
	fmt.Println("Changed the value of num using its pointer", *numAddr)

}