package main

import (
	"fmt"
	"slices"
)

func dcp1(s string) int {
	mtx := Must(NewMatrix(s))
	marks := make(map[Point]bool)
	score := 0

	mtx.Traverse(func(x, y int, ch string) bool {
		p := Point{x, y}
		if !marks[p] {
			trace := make([]Point, 0)
			expandToEnd(p, mtx, ch, &trace)

			for _, p := range trace {
				marks[p] = true
			}
			score += price(trace)
		}

		return false
	})

	return score
}

func price(trace []Point) int {
	area := len(trace)

	byX := make(map[int][]int)
	byY := make(map[int][]int)
	for _, p := range trace {
		x, y := p.X, p.Y
		byX[x] = append(byX[x], y)
		byY[y] = append(byY[y], x)
	}
	// fmt.Println(byX)
	// fmt.Println(byY)

	perimeter := 0
	for _, ys := range byX {
		perimeter += intersection(ys)
	}
	for _, xs := range byY {
		perimeter += intersection(xs)
	}

	fmt.Println(area, perimeter)
	return area * perimeter
}

func intersection(ns []int) int {
	// at least 2 intersections (in->out)
	count := 2
  slices.Sort(ns)
	for i := 1; i < len(ns); i++ {
		if ns[i]-ns[i-1] > 1 {
			// find gap
			// one more (out-in)
			count += 2
		}
	}

	return count
}

func expandToEnd(p Point, mtx *Matrix, ch string, trace *[]Point) {
	*trace = append(*trace, p)
	dir4 := []Point{{p.X - 1, p.Y}, {p.X + 1, p.Y}, {p.X, p.Y - 1}, {p.X, p.Y + 1}}
	for _, next := range dir4 {
		if c, err := mtx.PointAt(next); err == nil && c == ch && !slices.Contains(*trace, next) {
			expandToEnd(next, mtx, ch, trace)
		}
	}
}
