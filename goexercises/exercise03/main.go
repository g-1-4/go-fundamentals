package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	/*
		Multi-line Input (bufio)

		Problem
		Read a full address (with spaces)

		Print:
		Your address is: <address>

		Constraint:
		Must use bufio.NewReader
	*/

	//reading address using bufio
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your address:")
	input, _ := reader.ReadString('\n')
	//printing the output
	addr := strings.TrimSpace(input)
	fmt.Printf("Your address is: %s", addr)
}