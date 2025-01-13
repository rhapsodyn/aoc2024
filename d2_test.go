package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestD2P1(t *testing.T) {
	a := d2p1(`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`)
	assert.Equal(t, 2, a)
}

func TestD2P2(t *testing.T) {
	a := d2p2(`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`)
	assert.Equal(t, 4, a)
}
