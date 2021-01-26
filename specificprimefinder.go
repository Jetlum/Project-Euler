package main

import (
	"fmt"
	"math"
)

func main() {
	count := 0
	for i := 2; i <= 110000; i++ {
		// Checks if the returned value is true
		if IsPrime(i) == true {
			// Count the prime numbers
			count++
			// If count is equal to 10001 stop the loop
			if count == 10001 {
				fmt.Printf("The %v prime number is: %v", count, i)
				break
			}
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
