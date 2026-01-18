package main

import "fmt"

type Student struct {
	ID    int
	Name  string
	Marks int
}

func main() {
	/*
		Problem
		type Student struct {
			ID    int
			Name  string
			Marks int
		}

		Task
		Ask user for student ID
		Update marks if found
		Else print "Student not found"
	*/

	students := []Student{
		{1, "abc", 25},
		{2, "def", 23},
		{3, "ghi", 12},
		{4, "jkl", 29},
		{5, "mno", 18},
	}
	var name string
	var marks int
	fmt.Println("Enter the student name to update marks")
	fmt.Scan(&name)
	fmt.Println("Enter the marks to update")
	fmt.Scan(&marks)
	found := false
	for idx, i := range students {
		if i.Name == name {
			students[idx].Marks += marks
			found = true
		}
	}
	if !found {
		fmt.Println("Student not found")
	} else {
		fmt.Println("Update list")
		fmt.Println(students)
	}
}