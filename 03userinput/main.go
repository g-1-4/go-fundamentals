package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcomeMessage := "Welcome to user input"
	fmt.Println(welcomeMessage)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter you rating for our service:")

	//comma ok || comma err syntax
	rating, _ := reader.ReadString('\n')
	fmt.Println("Thanks for your rating:", rating)
	fmt.Printf("Type of rating is: %T\n", rating)

}