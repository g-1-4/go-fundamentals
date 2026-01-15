package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	/*
		   Problem

			Take user input:
			Name
			Age

			Print:
			Hello <name>, you are <age> years old

			Constraint:
			Use fmt.Scan
	*/

	//reading name using fmt.Scan
	var Name string
	fmt.Println("Enter your name:")
	fmt.Scan(&Name)

	//reading age using bufio
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your age:")
	input, _ := reader.ReadString('\n')
	//removing newline character from age input and converting to int64
	Age, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)

	//printing the output
	fmt.Printf("Hello %s, you are %d years old", Name, Age)
	

}