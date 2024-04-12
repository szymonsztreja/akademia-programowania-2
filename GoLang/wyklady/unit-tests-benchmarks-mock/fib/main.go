// https://dave.cheney.net/2013/06/09/writing-table-driven-tests-in-go
package main

import "fmt"

func main() {
	fmt.Printf("Fib(17)   = %d\n", Fib(17))
	fmt.Printf("FibNR(17) = %d\n", FibNR(17))
}

func Fib(n int) int {
	if n < 2 {
		return n
	}

	return Fib(n-1) + Fib(n-2)
}

func FibNR(n int) int {
	var (
		f1 = 1
		f2 = 0
	)

	for i := 0; i < n; i++ {
		f1, f2 = f2, f1+f2
	}

	return f2
}
