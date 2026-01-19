package main

import "fmt"

func main() {
	/*
		Problem:
			Create a function that prints:
				"Function started"

		Task:
			Use defer to print:
				"Function ended"

		Print order:
			Function started
			Function ended

		Constraint:
			Defer must be inside the function
	*/
	myfunc()

}

func myfunc() {
	defer fmt.Println("Ended")
	fmt.Println("Started")
}