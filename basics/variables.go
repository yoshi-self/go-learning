package main

import (
	"fmt"
)

func main() {
	Explicit()
	Implicit()
	Var()
	PackageVariable()
}

func Explicit() {
	// Explicit declaration
	var n int
	var x, y, z int
	var (
		a, b int
		name string
	)

	// Assign values
	n = 5
	x, y, z = 1, 2, 3
	a, b = x, y
	name = "John"
	fmt.Println(n, x, y, z, a, b, name)
}

func Implicit() {
	// Implicit declaration and assignment
	i := 1 // i is int
	b := true // b is bool
	f := 3.14 // f is float64
	s := "abc" // s is string
	// type inference from function return type
	n := one() // n is int
	fmt.Printf("i:%T, b:%T, f:%T, s:%T, n:%T\n", i, b, f, s, n)
}

func one() int {
	return 1
}

func Var() {
	// This is the same as a := 1
	var a = 1

	// var can make declaration block
	var (
		n = 1
		s = "string"
		b = true
	)

	fmt.Println(a, n, s, b)
}

var v = 100 // this can't be "v := 100"

func PackageVariable() {
	// v can be seen in anywhere inside this package
	fmt.Println(v)
}
