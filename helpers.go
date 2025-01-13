package main

import (
	"errors"
	"fmt"
	"strings"
)

func Must[T any](ret T, err error) T {
	if err != nil {
		panic(err)
	}

	return ret
}

func Must0(err error) {
	if err != nil {
		panic(err)
	}
}

func Must2[T1, T2 any](ret1 T1, ret2 T2, err error) (T1, T2) {
	if err != nil {
		panic(err)
	}

	return ret1, ret2
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

type Point struct {
	x int
	y int
}

type Vector struct {
	x int
	y int
}

type Matrix struct {
	raw string
	row int
	col int
}

func NewMatrix(s string) (*Matrix, error) {
	// cut leading & tailing `br`s
	s = strings.Trim(s, "\n")
	lines := strings.Split(s, "\n")
	row := len(lines)
	col := len(lines[0])

	var withOutBr strings.Builder
	for i, l := range lines {
		if len(l) != col {
			return nil, errors.New(fmt.Sprintf("line-%d width: %d not match with line-1 width: %d", i, len(l), col))
		}
		withOutBr.WriteString(l)
	}

	return &Matrix{
		raw: withOutBr.String(),
		row: row,
		col: col,
	}, nil
}

type TraverseHandler func(x int, y int, ch string)

func (m *Matrix) Traverse(handler TraverseHandler) {
	for y := 0; y < m.col; y++ {
		for x := 0; x < m.row; x++ {
			handler(x, y, string(m.raw[y*m.col+x]))
		}
	}
}

func (m *Matrix) At(x, y int) (string, error) {
	if x < 0 || y < 0 || x > m.col-1 || y > m.row-1 {
		return "", errors.New("OutOfBound")
	}
	return string(m.raw[y*m.col+x]), nil
}
