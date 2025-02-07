package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBP1(t *testing.T)  {
  assert.Equal(t, 7, dbp1("0 1 10 99 999", 1))
  assert.Equal(t, 22, dbp1("125 17", 6))
  assert.Equal(t, 55312, dbp1("125 17", 25))
}

func TestDBP2(t *testing.T)  {
  assert.Equal(t, 55312, dbp2("125 17", 25))
}
