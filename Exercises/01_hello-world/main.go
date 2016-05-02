package main

import "fmt"

func main() {
	//Print Hello world
	fmt.Println("Hello, World")

	//Print variable with your name
	name := "Phillip"
	fmt.Println("Hello, ", name)

	//Get user input and print
	var yourName string
	fmt.Print("Please enter your name: ")
	fmt.Scan(&yourName)
	fmt.Println("Hello ", yourName)

}
