package main

import "fmt"

func main() {
	var age int // variable declaration
	fmt.Println("My age is", age)

	age = 29 //assignment
	fmt.Println("My age is now", age)

	age = 54 //assignment
	fmt.Println("Finally, my age is", age)

	var name = "Marv" //type inference
	fmt.Println("Hi, my name is", name)

	// Multiple variable declaration
	var fname, lname = "Marvin", "Turner"
	fmt.Println("My full name is:", fname, lname)
}
