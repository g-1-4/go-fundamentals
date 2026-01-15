package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	/*
		    Problem

			Hardcode:
			username := "admin"
			password := "1234"

			Take input username & password

			Print:
			"Login successful" OR
			"Invalid credentials"

			Constraint:
			Exact match required
	*/
	// Hardcoded credentials
	// taking input for username
	fmt.Println("Enter you username")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	username := strings.TrimSpace(input)
	fmt.Println("You entered:", username)

	// taking input for password
	fmt.Println("Enter your password")
	input, _ = reader.ReadString('\n')
	password := strings.TrimSpace(input)
	fmt.Println("You entered:", strings.Repeat("*", len(password)))

	// authentication
	if username == "admin" && password == "1234" {
		fmt.Println("Login successful")
	} else {
		fmt.Println("Invalid credentials: Wrong username or password")
	}
}