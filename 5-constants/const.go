package main

import (
	"fmt"
)

func main() {
	// declare a single constant
	const a = 50
	// error, cannot reassign
	// a = 89
	fmt.Println(a)

	// declare a group of constants
	const (
		name    = "Diego"
		age     = 37
		country = "Belize"
	)
	fmt.Println(name, "from", country, "aged", age)

	// error to try to change the value of a const
	// country = "Panama"

	// value of a const must be known at compile time
	// cannot be assigned to a returned func value
	// because the call happens at runtime
	// error
	// const b = math.Sqrt(4)

}
