package main

import (
	"fmt"
	"math"
)

func bhaskaraFormula(a, b, c float64) (float64, float64) {
	// Calculate the discriminant (delta)
	delta := b*b - 4*a*c

	// Check if the discriminant is negative (complex roots)
	if delta < 0 {
		fmt.Println("Complex roots are not supported in this implementation.")
		return 0, 0
	}

	// Calculate the two roots using the Bhaskara formula
	root1 := (-b + math.Sqrt(delta)) / (2 * a)
	root2 := (-b - math.Sqrt(delta)) / (2 * a)

	return root1, root2
}

func main() {
	// Read coefficients from console
	var a, b, c float64
	fmt.Print("Enter coefficient a: ")
	fmt.Scan(&a)

	fmt.Print("Enter coefficient b: ")
	fmt.Scan(&b)

	fmt.Print("Enter coefficient c: ")
	fmt.Scan(&c)

	// Calculate roots using the Bhaskara formula
	root1, root2 := bhaskaraFormula(a, b, c)

	// Display the roots
	fmt.Printf("Root 1: %f\n", root1)
	fmt.Printf("Root 2: %f\n", root2)
}
