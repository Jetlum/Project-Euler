package main

import (
	"fmt"
)

func main() {
	sum := 0
	// Iterate to get the multiples of 3 or 5 below 1000
	for i := 999; i > 0; i-- {
		if i%3 == 0 || i%5 == 0 {
			sum += i
			fmt.Println(sum)
		}
	}
}
