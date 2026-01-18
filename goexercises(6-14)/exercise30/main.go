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
		Input a string.

		Task
		Count frequency of each character
		Ignore spaces

		Constraint
		Use map
	*/
	fmt.Println("Enter a string or sentence tocount frequency of each character")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	freq := make(map[rune]int)
	for _, v := range input {
		if v != ' '{
			freq[v]++
		}
	}
	for k,v := range freq {
		fmt.Printf("%c : %d\n", k, v)
	}
}