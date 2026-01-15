package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	/*

		Problem
			Record start time
			Ask user to press Enter
			Record end time

		Print duration in seconds

		Constraint:
			Use time.Now() and Sub
	*/
	startTime := time.Now()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press Enter to stop the timer...")
	_, _ = reader.ReadString('\n')
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("Duration: %.2f seconds", duration.Seconds())
}