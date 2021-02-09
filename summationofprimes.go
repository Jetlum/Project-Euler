package main

import (
	"fmt"
	"math"
)

func main() {
	sum := 0
	for i := 2; i < 2000000; i++ {
		// Checks if the returned value is true
		if IsPrime(i) == true {
			// Calculate the sum of all primes till two million
			sum += i
			fmt.Printf("Prime number %v and the sum is: %v\n", i, sum)
		}
	}

}

// IsPrime finds if a given number is a prime number
func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
