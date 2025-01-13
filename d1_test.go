package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestD1P1(t *testing.T) {
	ans := d1p1(`3   4
4   3
2   5
1   3
3   9
3   3`)
	assert.Equal(t, 11, ans)
}

func TestD1P2(t *testing.T) {
	ans := d1p2(`3   4
4   3
2   5
1   3
3   9
3   3`)
	assert.Equal(t, 31, ans)
}
