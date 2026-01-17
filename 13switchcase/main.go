package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and Case in Golang")

	// generarting random number between 1 to 6 for a die
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("The value of Dice is:", diceNumber)

	// switch case
	switch diceNumber{
	case 1:
		fmt.Println("Dice value is one and you can open")
	case 2:
		fmt.Println("you can move 2 spots")
	case 3:
		fmt.Println("you can move 3 spots")
		fallthrough
	case 4:
		fmt.Println("you can move 4 spots")
		fallthrough
	case 5:
		fmt.Println("you can move 5 spots")
	case 6:
		fmt.Println("you can move 6 spots and roll the die again")
	default:
		fmt.Println("Its Impossible!!!!")
	}
}