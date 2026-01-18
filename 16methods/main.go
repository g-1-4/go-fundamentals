package main

import "fmt"

type User struct {
	Name string
	Email string
	Status bool
	Age int
	oneAge int
}

func main() {
	fmt.Println("Methods in go language")

	u1 := User{
		Name: "Gowtham",
		Email: "gow@gmail.com",
		Status: true,
		Age: 21,
	}

	fmt.Println(u1)
	fmt.Printf("Name of the user is %+v\n", u1)
	fmt.Printf("Name of the user is %v, Email is %v\n", u1.Name, u1.Email)
	u1.getStatus()
	u1.newMail()
	fmt.Printf("Name of the user is %v, Email is %v\n", u1.Name, u1.Email)
}

func (u User) getStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u *User) newMail() {
	u.Email = "gowsri@gmail.com"
	fmt.Println("The new mail is:", u.Email)
}