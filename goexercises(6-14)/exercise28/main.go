package main

import "fmt"

func main() {
	/*
		Problem
		Given a slice of integers.

		Task
		Ask user for a number
		Remove first occurrence

		Print updated slice
		Constraint
		Use slicing
		If not found, print message
	*/
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Enter the number to be removed from the slice")
	var rnum int
	fmt.Scan(&rnum)
	flag := false
	for idx, ele := range nums {
		if ele == rnum {
			nums = append(nums[:idx], nums[idx+1:]...)
			flag = true
			break
		} 
	}
	if !flag {
		fmt.Println("Element not found in slice")
	} else{
		fmt.Println(nums)
	}
}