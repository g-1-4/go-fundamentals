package main

import "fmt"

type BankAccount struct {
	AC      int
	Name    string
	Balance int
}

func main() {
	/*
		Problem:
			Create a BankAccount struct with field:
				Balance (int)

		Task:
			Create a method to add money to balance

		Print:
			"Updated balance: <balance>"

		Constraint:
			Pointer receiver required
	*/
	u1 := BankAccount{
		AC:      1,
		Name:    "abc",
		Balance: 200,
	}
	u1.deposit(200)
	fmt.Println("Updated balance:", u1.Balance)
}

func (u *BankAccount) deposit(n int) {
	u.Balance += n
}