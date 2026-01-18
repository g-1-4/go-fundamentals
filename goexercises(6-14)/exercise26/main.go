package main

import "fmt"

type Product struct {
	ID       int
	Name     string
	Quantity int
}

func main() {
	/*
		Problem
		type Product struct {
			ID       int
			Name     string
			Quantity int
		}

		Task
		Store products in a slice
		Print products with quantity < 5
	*/

	p1 := Product{1, "tap", 3}
	p2 := Product{2, "faucet", 6}
	p3 := Product{3, "pipes", 1}
	p4 := Product{4, "gum", 2}
	p5 := Product{5, "filters", 5}

	myProducts := []Product{p1, p2, p3, p4, p5}

	fmt.Println("Products with quantity less than 5")
	for _, i := range myProducts {
		if i.Quantity < 5 {
			fmt.Printf("Id: %d, Name: %s, Qty: %d\n", i.ID, i.Name, i.Quantity)
		}
	}
}