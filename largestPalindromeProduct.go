package main

import "fmt"

func main() {
	// Declaring and initializing
	multiply, maxValue := 0, 0
	// Iterate through all three-digit numbers
	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			// Get the product of two three-digit numbers
			multiply = i * j
			// Checks if the palindrome is the same as the product of two three-digit numbers
			if ReverseNumber(multiply) == multiply {
				// Checks if the product of two three-digit numbers is bigger than the maxValue
				if multiply > maxValue {
					// Assign the biggest product of two three-digit numbers which is also the largest palindrome to the maxValue
					maxValue = multiply
					// Prints the largest palindrome
					fmt.Println("Max:", maxValue)
				}
			}
		}
	}
}

// ReverseNumber reverses the given number n
func ReverseNumber(n int) int {
	newInt := 0
	for n != 0 {
		remainder := n % 10
		newInt *= 10
		newInt += remainder
		n /= 10
	}
	return newInt
}
