package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers in go")

	// var ptr *int
	// fmt.Println("The value of pointer is", ptr)

	myNumber := 25
	var ptr *int = &myNumber
	fmt.Println("The number is:", myNumber)
	fmt.Println("The memory address of \"myNumber\" is", ptr)
	fmt.Println("The value at the memory address is", *ptr)

	*ptr -= 2
	fmt.Println("The value of myNumber after changing the value using pointer is:", myNumber)

}