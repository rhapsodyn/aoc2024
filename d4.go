package main

func d4p1(s string) int {
	m := Must(NewMatrix(s))
	counter := 0
	m.Traverse(func(x, y int, ch string) {
		if ch == "X" {
			counter += howManyXmas(m, x, y)
		}
	})
	return counter
}

var offsets8 [8][3]Vector = [8][3]Vector{
	{Vector{x: 1, y: 0}, Vector{x: 2, y: 0}, Vector{x: 3, y: 0}},       // to right
	{Vector{x: -1, y: 0}, Vector{x: -2, y: 0}, Vector{x: -3, y: 0}},    // to left
	{Vector{x: 1, y: -1}, Vector{x: 2, y: -2}, Vector{x: 3, y: -3}},    // to topRight
	{Vector{x: -1, y: -1}, Vector{x: -2, y: -2}, Vector{x: -3, y: -3}}, // to topLeft
	{Vector{x: 0, y: 1}, Vector{x: 0, y: 2}, Vector{x: 0, y: 3}},       // to top
	{Vector{x: 0, y: -1}, Vector{x: 0, y: -2}, Vector{x: 0, y: -3}},    // to bottom
	{Vector{x: 1, y: 1}, Vector{x: 2, y: 2}, Vector{x: 3, y: 3}},       // to bottomRight
	{Vector{x: -1, y: 1}, Vector{x: -2, y: 2}, Vector{x: -3, y: 3}},    // to bottomLeft
}

func howManyXmas(m *Matrix, x, y int) int {
	n := 0
	// 8 directions:
	for _, dir := range offsets8 {
		letter, err := m.At(x+dir[0].x, y+dir[0].y)
		if err != nil || letter != "M" {
			continue
		}
		letter, err = m.At(x+dir[1].x, y+dir[1].y)
		if err != nil || letter != "A" {
			continue
		}

		letter, err = m.At(x+dir[2].x, y+dir[2].y)
		if err == nil && letter == "S" {
			// found one
			n++
		}
	}
	return n
}

func d4p2(s string) int {
	counter := 0
	mx := Must(NewMatrix(s))
	mx.Traverse(func(x, y int, ch string) {
		if ch == "A" && crossfire(x, y, mx) {
			counter++
		}
	})
	return counter
}

func crossfire(x, y int, mx *Matrix) bool {
	tl, err := mx.At(x-1, y-1)
	if err != nil {
		return false
	}
	tr, err := mx.At(x+1, y-1)
	if err != nil {
		return false
	}
	bl, err := mx.At(x-1, y+1)
	if err != nil {
		return false
	}
	br, err := mx.At(x+1, y+1)
	if err != nil {
		return false
	}

	tlToBrOk := (tl == "M" && br == "S") || (tl == "S" && br == "M")
	trToBlOk := (tr == "M" && bl == "S") || (tr == "S" && bl == "M")
	return tlToBrOk && trToBlOk
}
