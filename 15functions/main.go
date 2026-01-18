package main

import "fmt"

func main() {

	greetUser()

	fmt.Println("Functions in golang")


	// func greeter()  {
	// 	fmt.Println("This is nested function")
	// } not allowed in go

	res := adder(2, 5)
	fmt.Println("Result is:", res)

	proRes := proAdder(1,2,3,4,5)
	fmt.Println("Result is:", proRes)
}

func adder(num1, num2 int) int  {
	return num1 + num2
}

func proAdder(values ...int) int {
	sum := 0
	for _, val := range values {
		sum += val
	}
	return sum
}

func greetUser() {
	fmt.Println("Welcome to the session")
}