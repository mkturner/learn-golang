package main

import "fmt"

func main() {
	// Go is very strict about explicit typing
	// no automatic promotion/ conversion

	i := 55
	j := 67.8

	// error
	// sum := i + j // int + float64 not allowed
	// convert j to int for matching types
	sum := i + int(j)
	fmt.Println(sum)

	// you must also use conversion to store a value
	// of one type in a variable of another type
	k := 10
	var l float64
	// error
	// l = k
	// must do explicit conversion
	l = float64(k)
	fmt.Println("l =", l)
}
