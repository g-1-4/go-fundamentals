package main

import "fmt"

type Rectange struct {
	Length int
	Width  int
}

func main() {
	/*
		Problem:
			Create a Rectangle struct with fields:
				Length, Width (int)

		Task:
			Create a method that returns area

		Print:
			"Area is: <value>"

		Constraint:
			Method must return value
	*/
	r1 := Rectange{
		Length: 8,
		Width:  7,
	}
	fmt.Println("Area is:", r1.Area())
}

func (r Rectange) Area() int {
	return r.Length * r.Width
}