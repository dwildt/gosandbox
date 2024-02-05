package main

import (
	"fmt"
)

func main() {
	// get a string from input and print each one of the characters inside the string
	fmt.Println("Enter a string: ")
	var input string
	fmt.Scanln(&input)
	for i := 0; i < len(input); i++ {
		fmt.Printf("%c\n", input[i])
	}

	// get the length of the string
	fmt.Println("len(input)")
	fmt.Println(len(input))

}
