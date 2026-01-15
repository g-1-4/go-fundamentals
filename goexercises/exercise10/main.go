package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		Problem
			Get current time
			Extract hour

		Print:
			0–11 → "Good Morning"
			12–16 → "Good Afternoon"
			17–23 → "Good Evening"

		Constraint:
			Use time.Now()
			Only if / else
	*/

	currentHour := time.Now().Hour()
	if(currentHour >=0 && currentHour <= 11) {
		fmt.Println("Good Morning")
	} else if(currentHour >=12 && currentHour <= 16) {
		fmt.Println("Good Afternoon")
	} else {
		fmt.Println("Good Evening")
	}
}