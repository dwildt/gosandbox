package main

import (
	"strconv"
)

func fizzbuzz(number int) string {

	if number == 0 {
		return "0"
	}

	if number < 0 {
		return "err"
	}

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
