package main

import "fmt"

func main() {
	/*
		Problem
		Candidates: A, B, C.

		Task
		Take votes using a loop
		Stop when user enters "exit"
		Print winner
	*/
	candidates := make(map[rune]int)
	candidates['A'] = 0
	candidates['B'] = 0
	candidates['C'] = 0
	totalVotes := 0
	maxVote := -1
	var winner rune

	for {
		var choice string

		fmt.Println("\nEnter A, B, C to vote or type exit to stop:")
		fmt.Scan(&choice)

		switch choice {
		case "A", "a":
			candidates['A']++
			totalVotes++
		case "B", "b":
			candidates['B']++
			totalVotes++
		case "C", "c":
			candidates['C']++
			totalVotes++
		case "exit":
			for k, v := range candidates {
				if v > maxVote {
					maxVote = v
					winner = k
				}
			}
			fmt.Printf("Total votes: %d\n", totalVotes)
			fmt.Printf("Winner is %c with %d votes\n", winner, maxVote)
			return
		default:
			fmt.Println("Invalid vote")
		}
	}
}