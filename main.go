package main

import "fmt"

// isAdmin :=true   // short form does not work outside function

func main() {
	// var name string
	// name = "Piash Islam"

	// var name string = "Piash Islam"

	// name := "Piash Islam"

	//grouped variables
	// var (
	// 	name string = "Piash Islma"
	// 	age  int    = 26
	// )

	//multiple variables declaration
	// var x, y int
	// x = 25
	// y = 30

	// var x, y int = 25, 30

	var x, y string = "Piash ", "Islam"

	isAdmin := true

	// var name = "Piash Islam"
	// fmt.Println(name)
	// fmt.Println(age)

	fmt.Println(x, y)
	fmt.Println(isAdmin)
}
