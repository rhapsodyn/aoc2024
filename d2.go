package main

import (
	"fmt"
	"strconv"
	"strings"
)

func d2parse(input string) [][]int {
	lines := strings.Split(input, "\n")
	var matrix [][]int
	for i, l := range lines {
		if len(l) > 0 {
			nums := strings.Fields(l)
			row := make([]int, len(nums))
			matrix = append(matrix, row)
			for j, n := range nums {
				matrix[i][j] = Must(strconv.Atoi(n))
			}
		}
	}

	return matrix
}

func safe(levels []int) bool {
	asc := levels[0] < levels[1]
	for i, n := range levels {
		if i < len(levels)-1 {
			if asc {
				if levels[i+1] <= n || levels[i+1]-n > 3 {
					return false
				}
			} else {
				// desc
				if levels[i+1] >= n || n-levels[i+1] > 3 {
					return false
				}
			}
		}
	}

	return true
}

func d2p1(s string) int {
	mat := d2parse(s)
	// fmt.Println(mat)
	n := 0
	for _, row := range mat {
		if safe(row) {
			fmt.Println(row)
			n++
		}
	}

	return n
}

func d2p2(s string) int {
	mat := d2parse(s)
	n := 0
	for _, row := range mat {
		if safe(row) {
			n++
		} else {
			for i := 0; i < len(row); i++ {
				var sub []int
				sub = append(sub, row[:i]...)
				sub = append(sub, row[i+1:]...)
				fmt.Println("sub: ", sub)
        if safe(sub) {
          n++
          break
        }
			}
		}
	}
	return n
}
