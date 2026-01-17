package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array in golang")

	// Creating an array
	var fruitList [4]string

	// Assigning values to array
	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	// fruitList[2] = "Banana"
	fruitList[3] = "Grapes"

	// Printing the array and its length
	fmt.Println("The fruitList is:", fruitList)
	fmt.Println("The length of fruitList is:", len(fruitList))

	// Another way to declare and initialize an array
	var vegList = [3]string{"Potato", "Tomato", "Cabbage"}

	// Printing the vegList and its length
	fmt.Println("The vegList is:", vegList)
	fmt.Println("The length of vegList is:", len(vegList))

	// Printing individual elements of the array
	fmt.Println("The first element of fruitList is:", fruitList[0])
	fmt.Printf("The last element of vegList is: %v\n", vegList[len(vegList)-1])
}