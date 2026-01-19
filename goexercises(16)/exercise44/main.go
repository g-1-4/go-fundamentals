package main

import "fmt"

type User struct {
	IsActive bool
}

func main() {
	/*
		Problem:
			Create a User struct with field:
				Active (bool)

		Task:
			Create a method to deactivate the user

		Print:
			"User deactivated"

		Constraint:
			Must use pointer receiver
	*/
	u1 := User{
		IsActive: true,
	}
	u1.deactivate()
	fmt.Println("Status of user:", u1.IsActive)
}

func (u *User) deactivate() {
	u.IsActive = false
	fmt.Println("user deactivated")
}