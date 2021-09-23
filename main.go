package main

import "fmt"

// Go support 3 types of datatypes

// ATOMIC Data type
// string
// int -> int32, int64
// uint -> uint32, uint64

// unsafe
// Pointers

// Abstract datatypes
// map[]<datatypes>
// struct{}
// interface{}

func main() {
	var a int
	a = 34
	fmt.Println("Value for a", a)

	var b = 45
	fmt.Println("Value for b", b)

	//  multiple initialization and assignments
	var (
		aa = 90
		bb = 34
	)
	fmt.Println("Value for aa+bb", aa+bb)

	// Type casting
	var x int16
	x = 4
	var y int32
	y = 8

	fmt.Println("Type casing in GO is ", int32(x)+y)

	// One line declaration
	m := 20
	n := 30

	fmt.Println("Sum of m and n :", m+n)

}
