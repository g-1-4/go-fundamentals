package main

import "fmt"

type Student struct {
	Name  string
	Marks int
}

func main() {
	/*
		Problem:
			Create a Student struct with fields:
				Name (string)
				Marks (int)

		Task:
			Create a method to check pass or fail

		Print:
			"Pass" OR "Fail"

		Constraint:
			Pass if marks >= 40
	*/
	s1 := Student{}
	s2 := Student{}
	s1.addStudent("abc", 45)
	s2.addStudent("def", 39)
	s1.checkRes()
	s2.checkRes()
}

func (s *Student) addStudent(name string, marks int) {
	s.Name = name
	s.Marks = marks
}

func (s Student) checkRes() {
	if s.Marks >= 40 {
		fmt.Printf("%v: Pass\n", s.Name)
	} else {
		fmt.Printf("%v: Fail\n", s.Name)
	}
}