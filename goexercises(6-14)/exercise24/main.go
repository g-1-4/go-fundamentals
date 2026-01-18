package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name  string
	Marks int
}

func main() {
	/*
		Problem
		type Student struct {
			Name  string
			Marks int
		}


		Task
		Store students in a map
		Find and print topper
	*/
	s1 := Student{"abc", 25}
	s2 := Student{"def", 23}
	s3 := Student{"ghi", 12}
	s4 := Student{"jkl", 29}
	s5 := Student{"mno", 18}

	studentMap := make(map[string]int)
	studentMap[s1.Name] = s1.Marks
	studentMap[s2.Name] = s2.Marks
	studentMap[s3.Name] = s3.Marks
	studentMap[s4.Name] = s4.Marks
	studentMap[s5.Name] = s5.Marks

	maxMarks := math.MinInt
	var studentName string

	for k,v := range studentMap {
		if v > maxMarks {
			studentName = k
			maxMarks = v
		}
	}

	fmt.Printf("The student with highest marks is %v", studentName)
}