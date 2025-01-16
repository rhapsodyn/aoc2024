package main

import (
	"slices"
	"strconv"
	"strings"
)

func d1parse(input string) (ls []int, rs []int) {
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		if len(l) > 0 {
			nums := strings.Fields(l)
			ls = append(ls, Must(strconv.Atoi(nums[0])))
			rs = append(rs, Must(strconv.Atoi(nums[1])))
		}
	}

	// fmt.Println(ls, rs)
	return
}

func d1p1(input string) int {
	ls, rs := d1parse(input)
	slices.SortFunc(ls, func(l, r int) int { return l - r })
	slices.SortFunc(rs, func(l, r int) int { return l - r })
	dis := 0
	for i, l := range ls {
		r := rs[i]
		dis += Abs(l - r)
	}

	return dis
}

func d1p2(s string) int {
	ls, rs := d1parse(s)
	rMap := make(map[int]int)
	for _, r := range rs {
		rMap[r]++
	}

	a := 0
	for _, l := range ls {
		a += l * rMap[l]
	}

	return a
}
