package main

import "fmt"

func main() {
	/*
		Problem
		Create a map of string -> int.

		Task
		Store 5 students and their marks
		Print all students and marks

		Constraint
		Use for range
	*/

	studentMap := make(map[string]int)
	studentMap["abc"] = 25
	studentMap["def"] = 18
	studentMap["ghi"] = 29
	studentMap["jkl"] = 14
	studentMap["mno"] = 22

	for k, v := range studentMap {
		fmt.Printf("Student %s scored %d marks for 30 marks\n", k, v)
	}

}