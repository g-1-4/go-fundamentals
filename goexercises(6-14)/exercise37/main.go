package main

import "fmt"

func main() {
	/*
		Problem
		Create a map of string -> int.

		Task
		Try modifying map value using pointer
		Observe and fix the issue
		Print map before and after

		Goal
		Understand why map values cannot be directly referenced
	*/
	myMap := make(map[string]int)
	myMap["gowtham"] = 1
	myMap["arepalli"] = 3
	myMap["sriram"] = 2

	ptr := &myMap

	// ptra := &myMap["gowtham"] not possible

	(*ptr)["gowtham"] = 2

	for k,v := range myMap {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}
}