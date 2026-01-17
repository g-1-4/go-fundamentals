package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in golang")

	// declaring a slice
	var fruitList = []string{"apple", "banana", "grapes"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	// adding more fruits to the slice
	fruitList = append(fruitList, "mango", "orange")
	fmt.Println("Fruit list is:", fruitList)

	//slicing the fruitList 1 to default
	fruitList = fruitList[1:]
	fmt.Println("Fruit list is:", fruitList)

	//slicing the fruitlist 1 to 3
	fruitList = fruitList[1:3]
	fmt.Println("Fruit list is:", fruitList)

	//slicing the fruitlist defualt to 3
	fruitList = fruitList[:3]
	fmt.Println("Fruit list is:", fruitList)


	// declaring slice using make
	highscores := make([]int, 4)
	
	// adding scores to the slice
	highscores[0] = 456
	highscores[1] = 646
	highscores[2] = 330
	highscores[3] = 986

	// trying to add 5th element
	// highscores[5] = 730
	highscores = append(highscores, 730, 240, 870)

	// printing the slice
	fmt.Println(highscores)

	// check if the slice is sorted
	fmt.Println(sort.IntsAreSorted(highscores))
	// sorting the slice
	sort.Ints(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))
	fmt.Println(highscores)


	// removing an element from slice

	var cities = []string{"hyderabad","chennai","pune","mumbai","delhi"}

	fmt.Println(cities)

	// remove pune from slice
	var index = 2

	cities = append(cities[0:index], cities[index+1:]...)

	fmt.Println(cities)

}