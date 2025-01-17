package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Equation struct {
	Lhs int
	Rhs []int
}

func d7p1(s string) int {
	equations := parse7(s)
	fmt.Println(equations)

	sum := 0
	for _, eq := range equations {
		if testEq(&eq) {
			sum += eq.Lhs
		}
	}

	return sum
}

func testEq(equation *Equation) bool {
	total := PowInt(2, len(equation.Rhs)-1)
	// fmt.Println(total)
	for i := 0; i < total; i++ {
		n := equation.Rhs[0]
		for j := 0; j < len(equation.Rhs)-1; j++ {
			if i&(1<<j) != 0 {
				n += equation.Rhs[j+1]
			} else {
				n *= equation.Rhs[j+1]
			}
		}
		// fmt.Println(n)
		if n == equation.Lhs {
			return true
		}
	}
	return false
}

func parse7(s string) []Equation {
	s = strings.Trim(s, "\n")
	lines := strings.Split(s, "\n")
	var equations []Equation
	for _, l := range lines {
		twoParts := strings.Split(l, ":")
		lhs := Must(strconv.Atoi(twoParts[0]))
		var rhs []int
		nums := strings.TrimSpace(twoParts[1])
		for _, n := range strings.Fields(nums) {
			rhs = append(rhs, Must(strconv.Atoi(n)))
		}

		equations = append(equations, Equation{Lhs: lhs, Rhs: rhs})
	}
	return equations
}

func d7p2(s string) int {
	equations := parse7(s)

	sum := 0
	for _, eq := range equations {
		if testEq2(&eq) {
			sum += eq.Lhs
		}
	}

	return sum
}

func testEq2(equation *Equation) bool {
	n := len(equation.Rhs) - 1
	total := PowInt(3, n)
	for i := 0; i < total; i++ {
		nth := i
		left := equation.Rhs[0]
		for j := 0; j < n; j++ {
			op := nth % 3
			left = calc7(left, equation.Rhs[j+1], op)
			nth /= 3
		}

		if left == equation.Lhs {
			return true
		}
	}

	return false
}

func calc7(left, right, op int) int {
	switch op {
	case 0:
		// add
		return left + right
	case 1:
		// mul
		return left * right
	case 2:
		// concat
		n := len(strconv.Itoa(right))
		times := PowInt(10, n)
		return left*times + right
	}

	panic("unreachable")
}
