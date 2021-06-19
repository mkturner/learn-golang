package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// numerics: signed/unsigned int, float, complex, byte, rune

	//Signed
	// int8: 8 bit signed integers
	// -128 to 127

	// int16: 16 bit signed integers
	// -32768 to 32767

	// int32: 32 bit signed integers
	// -2147483648 to 2147483647

	// int64: 64 bit signed integers
	// -9223372036854775808 to 9223372036854775807

	// int 32 or 64 bit depending on platform
	// default to int unless there is explicit need for others

	var a int = 89
	b := 95
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))
	fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b))

	// Unsigned
	// uint8: 8 bit signed integers
	// 0 to 255

	// uint16: 16 bit signed integers
	// 0 to 65535

	// uint32: 32 bit signed integers
	// 0 to 4294967295

	// uint64: 64 bit signed integers
	// 0 to 18446744073709551615

	// uint 32 or 64 bit depending on platform

	// Floats
	// float32: 32 bit floats
	// float64: 64 bit floats
	c, d := 5.67, 8.97
	fmt.Printf("type of c %T, d is %T\n", c, d)
	sum := c + d
	diff := c - d
	fmt.Println("sum:", sum, "diff:", diff)

	no1, no2 := 56, 89
	fmt.Println("no1 =", no1, "\tno2 =", no2)
	fmt.Println("sum:", no1+no2, "\tdiff:", no1-no2)

	// Complex
	// complex64: have float32 real and float32 i parts
	// complex128: have float64 real and float64 i parts

	// builtin function complex() used to construct complex nums
	// with real and imaginary parts
	// both the real and imaginary parts must be of same type

	// complex numbers may also be created with the shorthand syntax
	e := 6 + 7i
	fmt.Println("the comple number's value is:", e)
	complex1 := complex(5, 7)
	complex2 := 8 + 27i
	complexSum := complex1 + complex2
	fmt.Println("sum:", complexSum)
	complexProd := complex1 * complex2
	fmt.Println("product:", complexProd)

	// Other Numerics
	// byte == uint8
	// rune == int32
}
