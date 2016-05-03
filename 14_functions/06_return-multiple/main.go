package main

import "fmt"

func main() {
	fmt.Println(greet("Jane ", "Doe "))
}

func greet(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)

}

// It can also be written this way
//func greet(fname, lname string) (s string, x string) {
//	s = fmt.Sprint(fname, lname)
//	x = fmt.Sprint(lname, fname)
//	return
//}
