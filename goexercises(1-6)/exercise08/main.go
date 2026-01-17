package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	/*
		Problem
		Generate a secure random number between 0–9

		Print:
		Secure random number: X

		Constraint:
		Use crypto/rand
		No math/rand allowed here
	*/

	secureRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(10))
	fmt.Println("Secure random number:", secureRandomNumber.Int64())
}