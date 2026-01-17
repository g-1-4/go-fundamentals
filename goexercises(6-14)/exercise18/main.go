package main

import "fmt"

func main() {
	/*
		Problem
		Given a map of student names and marks.

		Task
		Ask user for student name
		Print marks if found
		Print "Student not found" if missing

		Constraint
		Use map lookup with if
	*/
	studentMap := make(map[string]int)
	studentMap["abc"] = 25
	studentMap["def"] = 18
	studentMap["ghi"] = 29
	studentMap["jkl"] = 14
	studentMap["mno"] = 22

	var studentName string
	fmt.Println("Enter your name")
	fmt.Scan(&studentName)
	value, err := studentMap[studentName]
	if !err {
		fmt.Println("Student not found")
	} else {
		fmt.Printf("Marks of the %v is %v", studentName, value)
	}
}