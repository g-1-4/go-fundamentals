package main

import (
	"fmt"
	"time"
)

func main() {
	/*Problem

	Print:

	"Hello, Go"

	"Welcome to Golang"

	Current date and time

	Constraint:

	Use separate variables for each message

	*/

	// declaring variables using var and string type
	var message1 string = "Hello, Go"
	// declaring variable ysing only var and type inference
	var message2 = "Welcome to Golang"
	// declaring variable using walrus operator and time package to get current date and time
	currentDateTime := time.Now().Format("02/01/2006 Monday 15:04")

	// printing the output using println(new line) and printf(format verb %v for default format)
	fmt.Println("First message:", message1)
	fmt.Printf("Second message: %v\n", message2)
	fmt.Println("Time now is:", currentDateTime)
}