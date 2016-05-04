package main

import "fmt"

// returns a func that returns a string
func makeGreeter() func() string {
	// anonymous function assigned to a variable is func expression
	return func() string {
		return "Hello, World"
	}
}

func main() {
	// makeGreeter returns the anon func and assigns it to greet
	greet := makeGreeter()
	// runs the anon func that was assigned to greet
	fmt.Println(greet())
	fmt.Printf("%T\n", greet)
}
