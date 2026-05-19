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

	// var x, y string = "Piash ", "Islam"

	// isAdmin := true

	// var name = "Piash Islam"
	// fmt.Println(name)
	// fmt.Println(age)

	// const pi = 3.14

	// fmt.Println(x, y, pi)
	// fmt.Println(isAdmin)

	// userName := "Piash"

	// fmt.Println(userName)

	// var name string = "GO Lang"
	// var isActive bool = true
	// var age int = 26
	// var floatingNumber = 20.25

	// fmt.Println(name)
	// fmt.Println(isActive)
	// fmt.Println(age)
	// fmt.Println(floatingNumber)

	// var age int
	// fmt.Println(age) //0

	// var name string
	// fmt.Println(name) //" "

	// var isAdmin bool
	// fmt.Println(isAdmin) //false

	// var score float64
	// fmt.Println(score) //0

	var name string = "next level"
	// age := 18
	rating := 4.5

	// fmt.Println("My name is", name)

	// fmt.Printf("My name is %s and I am %d years old with a rating of %.2f", name, age, rating)S

	formattedString := fmt.Sprintf("My name is %s and my rating is %.2f", name, rating)
	fmt.Println(formattedString)

}
