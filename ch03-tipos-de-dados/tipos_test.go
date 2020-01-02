package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestActive(t *testing.T) {
        var x, y int = 8, 6
	assert.Equal(t, soma(x,y), 14, "deveria ser 14")
}
