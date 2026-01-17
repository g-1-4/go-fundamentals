package main

import "fmt"

func main() {
	fmt.Println("Welcome to if/else in golang")

	loginCount := 17
	var res string
	if loginCount<10 {
		res = "Regular user"
	} else if loginCount>10{
		res = "Not a regular user"
	} else {
		res = "Exactly 10 messages"
	}

	fmt.Println(res)

	if 7%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is not less than 10")
	}
}