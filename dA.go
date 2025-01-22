package main

import (
	"strconv"
)

type PointSet9 map[Point]bool

func dap1(s string) int {
	mtx := Must(NewMatrix(s))
	total := 0
	mtx.Traverse(func(x, y int, ch string) bool {
		if ch == "0" {
			hs := make(PointSet9)
			waysTo9(x, y, mtx, 0, hs)
			total += len(hs)
		}
		return false
	})

	return total
}

func waysTo9(x, y int, m *Matrix, curr int, ps PointSet9) {
	if curr == 9 {
		ps[Point{x, y}] = true
	}

	left := Point{x - 1, y}
	right := Point{x + 1, y}
	top := Point{x, y - 1}
	bottom := Point{x, y + 1}
	dirs := [4]Point{left, right, top, bottom}

	next := curr + 1
	for _, d := range dirs {
		if ch, err := m.PointAt(d); err == nil && ch == strconv.Itoa(next) {
			waysTo9(d.X, d.Y, m, next, ps)
		}
	}
}

func dap2(s string) int {
	mtx := Must(NewMatrix(s))
	total := 0
	mtx.Traverse(func(x, y int, ch string) bool {
		if ch == "0" {
			prefixSet := make(map[string]bool)
			tracesTo9(x, y, mtx, 0, prefixSet, "")
			total += len(prefixSet)
		}
		return false
	})

	return total
}

func tracesTo9(x, y int, m *Matrix, curr int, ps map[string]bool, prefix string) {
	if curr == 9 {
		ps[prefix] = true
	}

	left := Point{x - 1, y}
	right := Point{x + 1, y}
	top := Point{x, y - 1}
	bottom := Point{x, y + 1}
	dirs := [4]Point{left, right, top, bottom}

	next := curr + 1
	for _, d := range dirs {
		if ch, err := m.PointAt(d); err == nil && ch == strconv.Itoa(next) {
			tracesTo9(d.X, d.Y, m, next, ps, prefix+d.String())
		}
	}
}
