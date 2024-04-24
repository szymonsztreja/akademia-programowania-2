package main

import "fmt"

func main() {
	fmt.Printf("FizzBuzz(3)  = %s\n", FizzBuzz(3))
	fmt.Printf("FizzBuzz(5)  = %s\n", FizzBuzz(5))
	fmt.Printf("FizzBuzz(15) = %s\n", FizzBuzz(15))
}

func FizzBuzz(n int) string {
	const (
		Fizz = "Fizz"
		Buzz = "Buzz"
	)

	var s string

	if n%3 == 0 {
		s += Fizz
	}

	if n%5 == 0 {
		s += Buzz
	}

	return s
}
