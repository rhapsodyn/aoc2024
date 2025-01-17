package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var i7 string = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func TestD7P1(t *testing.T) {
	assert.Equal(t, 3749, d7p1(i7))
}

func TestD7P2(t *testing.T) {
	assert.Equal(t, 11387, d7p2(i7))
}
