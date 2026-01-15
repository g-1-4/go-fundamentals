package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println("The present time is:", presentTime)

	// Formatting time
	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday"))

	// Creating a specific time
	createdDate := time.Date(2004, time.May, 5, 11, 11, 0, 0, time.UTC)
	fmt.Println("The created date is:", createdDate)
	fmt.Println("Formatted created data:", createdDate.Format("02/01/2006 Monday 15:04"))
}