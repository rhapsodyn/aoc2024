package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDCP1(t *testing.T) {
	assert.Equal(t, 140, dcp1(`AAAA
BBCD
BBCC
EEEC`))
	assert.Equal(t, 772, dcp1(`OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`))
	assert.Equal(t, 1930, dcp1(`RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`))
}
