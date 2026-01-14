package main

import "fmt"

var GlobalVariable string = "I am a global variable" // Global variable

func main() {
	var username string = "gowtham"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.455456456456
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values and aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Variable is of type: %T \n", anotherString)

	var anotherFloat float64
	fmt.Println(anotherFloat)
	fmt.Printf("Variable is of type: %T \n", anotherFloat)

	//implicit type
	var website = "google.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	//no var style
	numberOfUsers := 300
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type: %T \n", numberOfUsers)

	fmt.Println(GlobalVariable)
	fmt.Printf("Variable is of type: %T \n", GlobalVariable)
}