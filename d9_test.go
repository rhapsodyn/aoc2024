package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var i9 string = "2333133121414131402"

func TestD9P1(t *testing.T) {
	assert.Equal(t, 1928, d9p1(i9))
}

func TestD9P2(t *testing.T) {
	assert.Equal(t, 2858, d9p2(i9))
}
