package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64
	// use range to loop over a list or collection
	// range returns an index (or key depending on what you are looping thru
	// and a value, here we are using only the value and a blank identifier
	// in place of the index
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
