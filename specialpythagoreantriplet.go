package main

import (
	"fmt"
	"math"
)

func main() {
	for a := 1.0; a <= 1000; a++ {
		for b := 1.0; b <= 1000; b++ {
			for c := 1.0; c <= 1000; c++ {
				if a+b+c == 1000 && (math.Pow(a, 2)+math.Pow(b, 2) == math.Pow(c, 2)) {
					fmt.Printf("a is %v, b is %v, c is %v, a*b*c=", a, b, c)
					fmt.Println(a * b * c)
				}
			}
		}
	}
}
