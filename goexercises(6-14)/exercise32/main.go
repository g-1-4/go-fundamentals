package main

import "fmt"

func main() {
	/*
		Problem
		Create a 3×3 integer array.

		Task
		Print matrix
		Print sum of all elements
	*/
	matrix := make([][]int, 3)
	for i := 0; i < 3; i++ {
		matrix[i] = make([]int, 3)
	}

	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = 1
		}
	}

	fmt.Println(matrix)
	sum:=0
	for i := range matrix {
		for j := range matrix[i] {
			sum += matrix[i][j]
		}
	}
	fmt.Println("Sum of all elements in matrix is:", sum)
}