package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var ia string = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func TestDAP1(t *testing.T)  {
  assert.Equal(t, 36, dap1(ia))
}

func TestDAP2(t *testing.T)  {
  assert.Equal(t, 81, dap2(ia))
}
