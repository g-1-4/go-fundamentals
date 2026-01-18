package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary int
}

func main() {
	/*
		Problem
		type Employee struct {
			ID     int
			Name   string
			Salary int
		}

		Task
		Create employee
		Create pointer to struct
		Update salary using pointer
	*/
	e1 := Employee{
		ID:     1,
		Name:   "abc",
		Salary: 5000,
	}
	ptr := &e1

	ptr.Salary += 5000

	fmt.Printf("Id: %v, Name: %v, Salary; %v", ptr.ID, ptr.Name, ptr.Salary)
}