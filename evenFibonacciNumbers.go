package main

import (
	"fmt"
)

func main() {
	sum := 0
	// Iterate till a number to get the even terms whose values do not exceed four million,
	// find the sum of the even-valued terms
	for x := 0; x < 35; x++ {
		if fib(x)%2 == 0 {
			sum += fib(x)
			fmt.Printf("Fibonacci %v = %v\n", x, sum)
		}
	}
}

// fib calculates the fibonacci terms of a given number x
func fib(x int) int {
	if x == 0 {
		return 0
	} else if x <= 2 {
		return 1
	} else {
		return fib(x-2) + fib(x-1)
	}
}
