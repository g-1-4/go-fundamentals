package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	/*
		Problem:
			Create a User struct with fields:
				Name (string)
				Age (int)

		Task:
			Create a method that prints user details

		Print:
			"Name: <name>, Age: <age>"

		Constraint:
			Use value receiver
	*/
	u1 := User{
		Name: "gowtham",
		Age:  22,
	}
	u1.printDetails()
}

func (u User) printDetails() {
	fmt.Printf("Name: %v, Age: %v", u.Name, u.Age)
}