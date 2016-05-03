package main

import "fmt"

func main() {
	fmt.Println(greet("Jane", "Doe"))
}

// func func_name(arg1 type, arg2 type) (return_var type)
func greet(fname string, lname string) (s string) {
	s = fmt.Sprint(fname, " ", lname)
	return
}

// This is not as readable as the straight return shown in the previous example
