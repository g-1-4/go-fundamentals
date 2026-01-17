package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	/*
		Problem
		Input a sentence.

		Task
		Count frequency of each word using a map
		Print word and count

		Constraint
		Use strings.Fields
	*/

	fmt.Println("Enter a sentence to know the frequency")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)
	words := strings.Fields(input)
	freq := make(map[string]int)

	for _, word := range words {
		word = strings.Trim(word, ".!?,")
		freq[word] += 1
	}

	for k,v := range freq {
		fmt.Printf("%v : %v\n", k, v)
	}
}