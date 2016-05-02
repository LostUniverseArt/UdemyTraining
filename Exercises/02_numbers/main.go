package main

import "fmt"

func main() {
	//divideTwo()
	//printEven()
	//fizzBuzz()
	addNaturalNum()

}

func addNaturalNum() {
	total := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			total += i
		} else if i%5 == 0 {
			total += i
		}
	}
	fmt.Println(total)
}

func fizzBuzz() {
	for i := 0; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println(i, " -- FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, " -- Fizz")
		} else if i%5 == 0 {
			fmt.Println(i, " -- Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func printEven() {
	//Print the even numbers between 1 and 100
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	//return 0
}

func divideTwo() {
	//User enters numbers
	var numOne int
	var numTwo int
	fmt.Print("Please enter the first number: ")
	fmt.Scan(&numOne)
	fmt.Print("Please enter the second number: ")
	fmt.Scan(&numTwo)
	fmt.Println(numOne, "/", numTwo, "=", numOne/numTwo)
	//return 0
}
