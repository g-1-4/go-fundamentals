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
		Create 3 employees
		Store in a slice
		Print all employee details
	*/

	e1 := Employee{
		ID:     1,
		Name:   "abc",
		Salary: 12000,
	}
	e2 := Employee{2, "def", 15000}
	e3 := Employee{ID: 3, Name: "ghi", Salary: 18000}
	employees := []Employee{}
	employees = append(employees, e1, e2, e3)
	for _, emp := range employees {
		fmt.Printf("ID: %d, Name: %s, Salary: %d\n", emp.ID, emp.Name, emp.Salary)
	}
}