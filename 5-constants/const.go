package main

import "fmt"

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
}
