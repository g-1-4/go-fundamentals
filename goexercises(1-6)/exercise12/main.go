package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	/*
		Problem
		Generate a 4-digit random number

		Print:
		Temporary password: XXXX
		If password >= 5000 → "Strong"
		Else → "Weak"

		Constraint:
		Use random + if/else only
	*/
	randomNumber, _ := rand.Int(rand.Reader, big.NewInt(9999-1000+1))

	password := randomNumber.Int64() + 1000
	if password >= 5000 {
		fmt.Printf("Generated password is %d, it is Strong", password)
	} else {
		fmt.Printf("Generated password is %d, it is Weak", password)
	}
}