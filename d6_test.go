package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var i6 string = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestD6P1(t *testing.T) {
  fmt.Println(i6)
	assert.Equal(t, 41, d6p1(i6))
}

func TestD6P2(t *testing.T) {
  fmt.Println(i6)
	assert.Equal(t, 6, d6p2(i6))
}
