package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const abcd string = `ABCD
ABCD
ABCD
ABCD`

func TestMatrixStuff(t *testing.T) {
	m := Must(NewMatrix(abcd))
	assert.Equal(t, 4, m.col)
	assert.Equal(t, 4, m.row)
	var sb strings.Builder
	m.Traverse(func(x, y int, ch string) {
		sb.WriteString(ch)
		assert.Equal(t, Must(m.At(x, y)), ch)
	})

	assert.Equal(t, strings.ReplaceAll(abcd, "\n", ""), sb.String())
}
