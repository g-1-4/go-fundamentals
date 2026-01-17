package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps in golang")

	// creating a map

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"

	fmt.Println("List of all languages:", languages)
	fmt.Println("JS shorts for", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages after deleting key ruby:", languages)

	// looping through the map

	for key, value := range languages {
		fmt.Printf("For key %v, the value is %v\n", key, value)
	}
}