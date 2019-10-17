package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(number int) string {

	if number%15 == 0 {
		return "fizzbuzz"
	}

	if number%5 == 0 {
		return "buzz"
	}

	if number%3 == 0 {
		return "fizz"
	}

	return strconv.Itoa(number)
}

func main() {
	fmt.Println(fizzbuzz(3))
	fmt.Println(fizzbuzz(5))
	fmt.Println(fizzbuzz(15))
}
