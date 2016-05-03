package main

import "fmt"

func main() {
	fmt.Println(greet("Jane", "Doe"))
}

//
func greet(fname, lname string) string {
	// Sprint is string print, instead of sending the output
	// to the standard out we are passing a string back
	return fmt.Sprint(fname, " " , lname)
}
