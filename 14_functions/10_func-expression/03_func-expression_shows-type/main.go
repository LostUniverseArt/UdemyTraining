package main

import "fmt"

func main() {
	// anonymous function assigned to a variable is func expression
	greeting := func() {
		fmt.Println("Hello, World")
	}

	greeting()
	fmt.Printf("%T\n", greeting)
}
