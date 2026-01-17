package main

import "fmt"

func main() {
	fmt.Println("Structs in go language")

	u1 := User{
		Name: "Gowtham",
		Email: "gow@gmail.com",
		Status: true,
		Age: 21,
	}

	fmt.Println(u1)
	fmt.Printf("Name of the user is %+v\n", u1)
	fmt.Printf("Name of the user is %v, Email is %v\n", u1.Name, u1.Email)
}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}