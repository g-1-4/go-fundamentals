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

		Input age

		Output:
		< 18 → "Minor"
		18–59 → "Adult"
		>= 60 → "Senior Citizen"

		Constraint:
		Only if / else if / else
	*/

	fmt.Println("Enter your age:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	age, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)

	if(age<18){
		fmt.Println("Minor")
	} else if(age>18 && age<=59){
		fmt.Printf("Adult\n")
	} else{
		fmt.Println("Senior Citizen")
	}
}