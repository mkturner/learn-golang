package main

import (
	"fmt"
	"math"
)

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

	// declaring with initial value
	var brosAge int = 8
	fmt.Println("My brother's age is", brosAge)

	// could also be written as:
	// var age = 8
	// when initial value present, type may be removed

	// multiple value declaration
	var width, height int = 100, 50
	fmt.Println("width is:", width, "height is:", height, "area is:", (width * height))

	// declaring variables of different types
	var (
		myName   = "marv"
		myAge    = 29
		myHeight int
	)

	fmt.Println("my name is", myName, "age is", myAge, "and height is", myHeight)

	// shorthand declaration :=
	// will infer type in functions/blocks
	count := 10
	fmt.Println("Count = ", count)

	// multiple shorthand declarations
	favoriteColor, favoriteFood := "gold", "pizza"
	fmt.Println("My favorite color is", favoriteColor, "and my favorite food is", favoriteFood)

	// error: each left-hand variable must have value assigned on right
	// name, age := "marvin"

	// at least one of the variables used in a shorthand syntax must be new

	// variables computed at runtime
	a, b := 145.8, 543.6
	c := math.Min(a, b)
	fmt.Println("Minimum value is", c)

	// REMEMBER: Go is strongly typed, vars of one type cannot be reassigned to another type
}
