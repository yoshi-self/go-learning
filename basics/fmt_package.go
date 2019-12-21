package main

import (
	"fmt"
)

func main() {
	// Print a string with a new line
	fmt.Println("Hello")

	// Takes multiple arguments, joins with spaces
	fmt.Println("I", "am", "awesome")

	// printf
	fmt.Printf("num: %d\n", 5)

	// shows "MISSING"
	fmt.Printf("month/day: %d/%d\n", 12)

	// shows "EXTRA"
	fmt.Printf("month/day: %d/%d\n", 12, 25, 0)

	fmt.Printf("\n")

	// %v supports variety types
	fmt.Printf("num=%v, string=%v, array=%v\n", 1, "Hello", [...]int{1, 2, 3})

	// %T shows type info
	fmt.Printf("num=%T, string=%T, array=%T\n", 1, "Hello", [...]int{1, 2, 3})

	// builtin print
	print("print\n")
	// builtin println
	println("print")
	println(1, 2, 3) // => "1 2 3"
	// NOTE: These output to stderr while fmt.Println outputs to stdout

}
