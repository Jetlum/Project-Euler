package main

import (
	"fmt"
	"math"
)

func main() {
	// Declaring variables
	var i, sumSquares, squareSum float64

	// Iterating first 100 natural numbers and finding the sum of the squares
	for i = 1; i <= 100; i++ {
		// fmt.Printf("%v: %v\n", i, math.Pow(i, 2))
		sumSquares += math.Pow(i, 2)
	}
	// Iterating first 100 natural numbers and finding the square of the sum
	for i = 1; i <= 100; i++ {
		squareSum += i
		// fmt.Println(squareSum)
	}

	fmt.Println("The difference between the sum of the squares of the first one hundred natural numbers and the square of the sum: ", (math.Pow(squareSum, 2) - sumSquares))
}
