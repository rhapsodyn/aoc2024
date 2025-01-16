package main

func d4p1(s string) int {
	m := Must(NewMatrix(s))
	counter := 0
	m.Traverse(func(x, y int, ch string) bool {
		if ch == "X" {
			counter += howManyXmas(m, x, y)
		}
    return false
	})
	return counter
}

var offsets8 [8][3]Vector = [8][3]Vector{
	{Vector{X: 1, Y: 0}, Vector{X: 2, Y: 0}, Vector{X: 3, Y: 0}},       // to right
	{Vector{X: -1, Y: 0}, Vector{X: -2, Y: 0}, Vector{X: -3, Y: 0}},    // to left
	{Vector{X: 1, Y: -1}, Vector{X: 2, Y: -2}, Vector{X: 3, Y: -3}},    // to topRight
	{Vector{X: -1, Y: -1}, Vector{X: -2, Y: -2}, Vector{X: -3, Y: -3}}, // to topLeft
	{Vector{X: 0, Y: 1}, Vector{X: 0, Y: 2}, Vector{X: 0, Y: 3}},       // to top
	{Vector{X: 0, Y: -1}, Vector{X: 0, Y: -2}, Vector{X: 0, Y: -3}},    // to bottom
	{Vector{X: 1, Y: 1}, Vector{X: 2, Y: 2}, Vector{X: 3, Y: 3}},       // to bottomRight
	{Vector{X: -1, Y: 1}, Vector{X: -2, Y: 2}, Vector{X: -3, Y: 3}},    // to bottomLeft
}

func howManyXmas(m *Matrix, x, y int) int {
	n := 0
	// 8 directions:
	for _, dir := range offsets8 {
		letter, err := m.At(x+dir[0].X, y+dir[0].Y)
		if err != nil || letter != "M" {
			continue
		}
		letter, err = m.At(x+dir[1].X, y+dir[1].Y)
		if err != nil || letter != "A" {
			continue
		}

		letter, err = m.At(x+dir[2].X, y+dir[2].Y)
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
	mx.Traverse(func(x, y int, ch string) bool {
		if ch == "A" && crossfire(x, y, mx) {
			counter++
		}
		return false
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
