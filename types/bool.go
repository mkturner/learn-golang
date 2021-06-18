package main

import "fmt"

func main() {
	// 3 major types: bool, numeric, string

	// bool: true | false
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)

	// logical AND operator
	c := a && b
	fmt.Println("c:", c)

	// logical OR operator
	d := a || b
	fmt.Println("d:", d)

}
