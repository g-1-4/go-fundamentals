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
		Given a slice of employees.

		Task
		Print employees with salary > 50000

		Constraint
		Use if inside for
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
		if emp.Salary >= 15000 {
			fmt.Printf("ID: %v, Name: %v\n", emp.ID, emp.Name)
		}
	}
}