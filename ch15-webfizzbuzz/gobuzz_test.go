package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestBuzz(t *testing.T) {
	assert.Equal(t, fizzbuzz(3), "fizz", "3 é fizz")
	assert.Equal(t, fizzbuzz(5), "buzz", "5 é buzz")
	assert.Equal(t, fizzbuzz(15), "fizzbuzz", "15 é fizzbuzz")

	assert.Equal(t, fizzbuzz(0), "0", "0 é 0")
	assert.Equal(t, fizzbuzz(4), "4", "4 é 4")

	assert.Equal(t, fizzbuzz(-2), "err", "-2 é err")

}
