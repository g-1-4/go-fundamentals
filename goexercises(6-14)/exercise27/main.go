package main

import "fmt"

func main() {
	/*
		Problem
		Given a slice of integers.

		Task
		Print duplicate values using a map
	*/

	nums := []int{1, 2, 1, 3, 2, 4, 4, 5, 6, 7, 1, 2, 2, 8, 9}

	numsMap := make(map[int]int)

	for _, num := range nums {
		numsMap[num] += 1
	}
	fmt.Println("Duplicates in this array is")
	for k, v := range numsMap {
		if v > 1 {
			fmt.Printf("%d occurs %d times\n", k, v)
		}
	}
}