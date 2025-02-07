package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type CacheB = map[int][2]int

func dbp1(s string, n int) int {
	s = strings.Trim(s, "\n")
	numstr := strings.Fields(s)
	var nums []int
	for _, ns := range numstr {
		nums = append(nums, Must(strconv.Atoi(ns)))
	}

	cb := make(CacheB)
	for i := 0; i < n; i++ {
		var newNums []int
		for _, nu := range nums {
			l, r := blinkGPT(nu, cb)
			newNums = append(newNums, l)
			if r != -1 {
				newNums = append(newNums, r)
			}
		}
		nums = newNums
		fmt.Println("i:", i, "cahced:", len(cb))
	}

	fmt.Println(nums)
	return len(nums)
}

// Log10 approach by chatgpt, though not helping
func blinkGPT(n int, cb CacheB) (int, int) {
	if arr, exits := cb[n]; exits {
		return arr[0], arr[1]
	}

	// Find the number of digits in the number
	digits := int(math.Floor(math.Log10(float64(n))) + 1)
	if n == 0 {
		return 1, -1
	} else if digits%2 != 0 {
		return n * 2024, -1
	} else {
		// Calculate the divisor to split the number
		divisor := int(math.Pow(10, float64(digits/2)))

		// Split the number into two parts
		left := n / divisor
		right := n % divisor
		cb[n] = [2]int{left, right}
		return left, right
	}
}

func dbp2(s string, n int) int {
	s = strings.Trim(s, "\n")
	numstr := strings.Fields(s)
	nums := make(map[int]int)
	for _, ns := range numstr {
		n := Must(strconv.Atoi(ns))
		nums[n] = 1
	}

	for i := 0; i < n; i++ {
		newNums := make(map[int]int)
		for n, v := range nums {
			l, r := blink(n)
			if _, exists := newNums[l]; exists {
				newNums[l] += v
			} else {
				newNums[l] = v
			}

			if r != -1 {
				if _, exists := newNums[r]; exists {
					newNums[r] += v
				} else {
					newNums[r] = v
				}
			}
		}
		nums = newNums
	}

	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

// r could be -1
func blink(n int) (int, int) {
	s := strconv.Itoa(n)
	l := len(s)

	if n == 0 {
		return 1, -1
	} else if l%2 != 0 {
		return n * 2024, -1
	} else {
		half := l / 2
		ls := s[:half]
		rs := s[half:]
		l := Must(strconv.Atoi(ls))
		r := Must(strconv.Atoi(rs))
		return l, r
	}
}
