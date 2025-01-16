package main

func d6p1(s string) int {
	mapp := Must(NewMatrix(s))

	var p Point
	mapp.Traverse(func(x, y int, ch string) bool {
		if ch == "^" {
			p = Point{x, y}
			return true
		} else {
			return false
		}
	})

	route, _ := findRoute(p, mapp)
	noDir := make(map[Point]bool)
	for k := range route {
		noDir[k.Point] = true
	}

	return len(noDir)
}

type Arrow struct {
	Point
	Face
}

func findRoute(p Point, mapp *Matrix) (map[Arrow]struct{}, bool) {
	route := make(map[Arrow]struct{})
	curr := Arrow{p, Up}

	for {
		// fmt.Printf("%d,%d,%s\n", curr.x, curr.y, Display(curr.Face))
		if _, exists := route[curr]; exists {
			return nil, true
		}

		// fmt.Println(cursor)
		route[curr] = struct{}{}
		next, oob := nextStep(curr, mapp)
		if oob {
			break
		}

		// go on
		curr = next
	}

	return route, false
}

func nextStep(curr Arrow, mapp *Matrix) (Arrow, bool) {
	next := Arrow{curr.Point, curr.Face}
	switch curr.Face {
	case Up:
		next.Y--
	case Down:
		next.Y++
	case Left:
		next.X--
	case Right:
		next.X++
	}

	ch, err := mapp.At(next.X, next.Y)
	if err != nil {
		// out of bound
		// fmt.Println("oob", next)
		return next, true
	}

	if ch == "#" {
		// block, turn right
		switch curr.Face {
		case Up:
			next.Face = Right
			next.X = curr.X + 1
			next.Y = curr.Y
		case Down:
			next.Face = Left
			next.X = curr.X - 1
			next.Y = curr.Y
		case Left:
			next.Face = Up
			next.X = curr.X
			next.Y = curr.Y - 1
		case Right:
			next.Face = Down
			next.X = curr.X
			next.Y = curr.Y + 1
		}
	}
	return next, false
}

func d6p2(s string) int {
	mapp := Must(NewMatrix(s))

	var starter Point
	mapp.Traverse(func(x, y int, ch string) bool {
		if ch == "^" {
			starter = Point{x, y}
			return true
		} else {
			return false
		}
	})
	// fmt.Println(starter)

	route, _ := findRoute(starter, mapp)
	noDir := make(map[Point]struct{})
	for k := range route {
		noDir[k.Point] = struct{}{}
	}
	delete(noDir, starter)

	counter := 0
	var info []Point
	for pos := range noDir {
		m := mapp.Clone()
		m.Set(pos.X, pos.Y, '#')
		if _, loop := findRoute(starter, &m); loop {
			// m.Pretty()
			// fmt.Println("loop")
			info = append(info, pos)
			counter++
		}
	}

	return counter
}
