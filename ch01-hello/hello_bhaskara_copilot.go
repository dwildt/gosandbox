package main

import (
	"fmt"
	"math"
)

func main() {
	// I need to calculate bhaskara's formula. Receive
	// the values of a, b and c and calculate the delta
	// and the roots of the equation
	// a = 1, b = -3, c = 2 >>> Expected output: Root 1: 2, Root 2: 1
	fmt.Println("Enter the values of a, b and c: ")
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	delta := b*b - 4*a*c
	fmt.Println("Delta: ", delta)
	if delta < 0 {
		fmt.Println("No real roots")
	} else {
		root1 := (-b + math.Sqrt(delta)) / (2 * a)
		root2 := (-b - math.Sqrt(delta)) / (2 * a)
		fmt.Println("Root 1: ", root1)
		fmt.Println("Root 2: ", root2)
	}
}
