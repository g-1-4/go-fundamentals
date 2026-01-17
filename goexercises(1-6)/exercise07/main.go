package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		Problem
		Generate a random number between 1 and 10

		Print:
		The number
		Whether it is greater than 5 or less than or equal to 5

		constraint:
		Use math/rand and time.Now().UnixNano()
	*/

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(10)+1
	println("The number:", randomNumber)
	if(randomNumber > 5) {
		fmt.Println("The number is greater than 5")
	}else {
		fmt.Println("The number is less than or equal to 5")
	}
}