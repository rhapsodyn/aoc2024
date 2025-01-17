package main

func d8p1(s string) int {
	mtx := MustMatrix(s)
	antennas := make(map[string][]Point)
	mtx.Traverse(func(x, y int, ch string) bool {
		p := Point{x, y}
		if ch != "." {
			if _, exists := antennas[ch]; !exists {
				antennas[ch] = make([]Point, 0)
			}

			antennas[ch] = append(antennas[ch], p)
		}

		return false
	})

	nodes := make(map[Point]bool)
	for _, v := range antennas {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				p1, p2 := findAntiNodes(v[i], v[j])
				if _, err := mtx.PointAt(p1); err == nil {
					nodes[p1] = true
				}
				if _, err := mtx.PointAt(p2); err == nil {
					nodes[p2] = true
				}
			}
		}
	}

	return len(nodes)
}

func findAntiNodes(point1, point2 Point) (Point, Point) {
	deltaX := point2.X - point1.X
	deltaY := point2.Y - point1.Y
	p1 := Point{X: point1.X - deltaX, Y: point1.Y - deltaY}
	p2 := Point{X: point2.X + deltaX, Y: point2.Y + deltaY}
	return p1, p2
}

func d8p2(s string) int {
	mtx := MustMatrix(s)
	antennas := make(map[string][]Point)
	mtx.Traverse(func(x, y int, ch string) bool {
		p := Point{x, y}
		if ch != "." {
			if _, exists := antennas[ch]; !exists {
				antennas[ch] = make([]Point, 0)
			}

			antennas[ch] = append(antennas[ch], p)
		}

		return false
	})

	nodes := make(map[Point]bool)
	for _, v := range antennas {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				ps := findAntiNodes2(v[i], v[j], mtx)
				for _, p := range ps {
					nodes[p] = true
				}
			}
		}
	}

	return len(nodes)
}

func findAntiNodes2(point1, point2 Point, mtx *Matrix) []Point {
	deltaX := point2.X - point1.X
	deltaY := point2.Y - point1.Y
	ps := make([]Point, 0)
	i := 0
	for {
		p := Point{point1.X - i*deltaX, point1.Y - i*deltaY}
		if _, err := mtx.PointAt(p); err != nil {
			break
		}
		ps = append(ps, p)
    i++
	}
	i = 0
	for {
		p := Point{point2.X + i*deltaX, point2.Y + i*deltaY}
		if _, err := mtx.PointAt(p); err != nil {
			break
		}
		ps = append(ps, p)
    i++
	}

	return ps
}
