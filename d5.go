package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Rule [2]int
type Pages []int

func d5parse(s string) ([]Rule, []Pages) {
	s = strings.Trim(s, "\n")
	lines := strings.Split(s, "\n")
	var rules []Rule
	var pagesList []Pages
	brk := false
	for _, l := range lines {
		if len(l) == 0 {
			brk = true
			continue
		}

		if !brk {
			// parsing rule
			two := strings.Split(l, "|")
			var rule [2]int
			rule[0] = Must(strconv.Atoi(two[0]))
			rule[1] = Must(strconv.Atoi(two[1]))

			rules = append(rules, rule)
		} else {
			// parsig pages
			nums := strings.Split(l, ",")
			var pages []int
			for _, n := range nums {
				pages = append(pages, Must(strconv.Atoi(n)))
			}

			pagesList = append(pagesList, pages)
		}
	}

	return rules, pagesList
}

func d5p1(s string) int {
	rules, pagess := d5parse(s)
	beforeAfters := make(map[int][]int)
	for _, r := range rules {
		if beforeAfters[r[0]] == nil {
			beforeAfters[r[0]] = make([]int, 0)
		}

		beforeAfters[r[0]] = append(beforeAfters[r[0]], r[1])
	}

	fmt.Println(beforeAfters)
	// fmt.Println(pagess)

	score := 0
	for _, pages := range pagess {
		ok := true
		for i := 0; i < len(pages)-1; i++ {
			before := pages[i]
			for j := i + 1; j < len(pages); j++ {
				after := pages[j]
				if !slices.Contains(beforeAfters[before], after) {
					ok = false
					break
				}
			}

			if !ok {
				break
			}
		}

		if ok {
			score += pages[(len(pages)-1)/2]
		}
	}

	return score
}

func d5p2(s string) int {
	rules, pagess := d5parse(s)
	beforeAfters := make(map[int][]int)
	keys := make(map[int]int)
	for _, r := range rules {
		if beforeAfters[r[0]] == nil {
			beforeAfters[r[0]] = make([]int, 0)
			keys[r[0]] = 0
		}

		beforeAfters[r[0]] = append(beforeAfters[r[0]], r[1])
	}

	// fmt.Println(len(beforeAfters))
	// uniqVals := make(map[int]bool)
	// for k, v := range beforeAfters {
	// 	fmt.Printf("%d:%v\n", k, v)
	// foo := 0
	// for _, val := range v {
	// 	keys[val]++
	// }
	// fmt.Println("hit count:", foo)
	// }
	// fmt.Println(uniqVals, len(uniqVals))
	// fmt.Println(keys)
	// return 0

	// fmt.Println(beforeAfters)
	// fmt.Println(pagess)

	score := 0
	for _, pages := range pagess {
		ok := true
		for i := 0; i < len(pages)-1; i++ {
			before := pages[i]
			for j := i + 1; j < len(pages); j++ {
				after := pages[j]
				if !slices.Contains(beforeAfters[before], after) {
					ok = false
					break
				}
			}

			if !ok {
				break
			}
		}

		if !ok {
			newOrder := resort(pages, beforeAfters)
			fmt.Println(newOrder)
			score += newOrder[(len(newOrder)-1)/2]
		}
	}

	return score
}

func resort(pages []int, beforeAfters map[int][]int) []int {
	// newOrder := make([]int, len(pages))
	ok := false
	for !ok {
		ok = true
		for i := 0; i < len(pages)-1; i++ {
			before := pages[i]
			for j := i + 1; j < len(pages); j++ {
				after := pages[j]
				if !slices.Contains(beforeAfters[before], after) {
					ok = false
					var newPages []int
					newPages = append(newPages, pages[:i]...)
					newPages = append(newPages, after)
					newPages = append(newPages, pages[i:j]...)
					newPages = append(newPages, pages[j+1:]...)

					// fmt.Println(pages)
					// fmt.Println(i, j)
					// fmt.Println(newPages)
					//
					// another iter
          pages = newPages
					break
				}
			}
		}
	}
	return pages
}
