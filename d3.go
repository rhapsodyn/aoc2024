package main

import (
	"regexp"
	"strconv"
	"strings"
)

func doMul(s string) int {
	mid := strings.Index(s, ",")
	prefixLen := len("mul(")
	suffixIdx := len(s) - 1
	l := Must(strconv.Atoi(s[prefixLen:mid]))
	r := Must(strconv.Atoi(s[mid+1 : suffixIdx]))
	return l * r
}

func d3p1(s string) int {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := re.FindAllStringSubmatch(s, -1)
	a := 0
	for _, m := range matches {
		a += doMul(m[0])
	}
	return a
}

func d3p2(s string) int {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	matches := re.FindAllStringSubmatch(s, -1)
	a := 0
	enable := true
	for _, m := range matches {
		if m[0] == "do()" {
			enable = true
		} else if m[0] == "don't()" {
			enable = false
		} else {
			if enable {
				a += doMul(m[0])
			}
		}
	}
	return a
}
