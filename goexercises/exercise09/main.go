package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		Problem
			Generate random number between 1–5
			Ask user to enter a number

		Print:
			"Correct guess" OR
			"Wrong guess"

		📌 Constraint:
			Only one attempt
			No loops
	*/

	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(5)+1

	var userGuess int
	fmt.Println("Enter a number between 1 and 5:")
	fmt.Scan(&userGuess)

	if(userGuess == randomNumber) {
		fmt.Println("Correct guess")
	} else {
		fmt.Println("Wrong guess, the correct number was", randomNumber)
	}
}