package main

import "fmt"

func main() {
	/*
		Problem
		Given a slice of integers.

		Task
		Count even numbers
		Count odd numbers
		Print both counts
	*/
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenCount := 0
	oddCount := 0

	for _, i := range mySlice {
		if i%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}
	fmt.Println("Count of Even numbers:", evenCount)
	fmt.Println("Count of Odd numbers:", oddCount)
	fmt.Println("Total number of elements:", len(mySlice))
}