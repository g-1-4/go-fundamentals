package main

import "fmt"

func main() {
	fmt.Println("Loops in go lang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}
	
	for idx, value := range days{
		fmt.Printf("Index is %v, and its value is %v\n", idx, value)
	}

	val := 1

	for val < 10 {

		if val == 5 {
			val++
			continue
		}

		if val == 7 {
			goto label1
		}

		fmt.Println(val)
		val++
	}

	label1:
		fmt.Println("Jumped from loop to this statement")

}