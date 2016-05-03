package main

import "fmt"

func main() {
	greet("Jane", "Doe")
}

// as both parameters are strings we can declare once
func greet(fname, lname string) {
	fmt.Println(fname, lname)

}
