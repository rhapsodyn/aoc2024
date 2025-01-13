package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestD3P1(t *testing.T) {
	a := d3p1("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
	assert.Equal(t, 161, a)
}

func TestD3P2(t *testing.T) {
	a := d3p2("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))")
	assert.Equal(t, 48, a)
}
