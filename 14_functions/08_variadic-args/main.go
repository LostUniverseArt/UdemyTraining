package main

import "fmt"

func main() {
	// [] denotes a slice, which is a list of items
	// adding ... to the end of the param tells Go to
	// open up the data variable and send each item
	// in the list one by one because the average func
	// requires float64 and the data var is a slice of
	// float64 they must be sent individually
	data := []float64{43, 56, 87, 12, 45, 57}
	n := average(data...)
	fmt.Println(n)
}

//
func average(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

//09_slice-param-arg
// This could also have been done this way
//func main() {
//	data := []float64{43, 56, 87, 12, 45, 57}
//	n := average(data)
//	fmt.Println(n)
//}
//
// receives a []float64 instead of a variadic
//func average(sf []float64) float64 {
//	total := 0.0
//	for _, v := range sf {
//		total += v
//	}
//	return total / float64(len(sf))
//}
