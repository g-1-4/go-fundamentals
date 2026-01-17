package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		Problem 1: Array Maximum

		Problem
		Create an array of 5 integers.

		Task
		Print all elements using a for loop
		Find and print the largest number

		Constraint
		Use only array indexing and for
	*/

	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	max := math.MinInt
	for i :=0 ; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	fmt.Println("The maximum element i this array is:", max)
}