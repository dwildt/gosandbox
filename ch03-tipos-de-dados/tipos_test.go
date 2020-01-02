package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestSoma(t *testing.T) {
        var x, y int = 8, 6
	assert.Equal(t, soma(x,y), 14, "deveria ser 14")
}

func TestNumeroPar(t *testing.T) {
    var p int32 = 25000
    var i byte = 3

    assert.Equal(t, par(p), true, "deveria ser true... 25000 é par")
    assert.Equal(t, par(int32(i)), false,"deveria ser falso... 3 é impar")
}

