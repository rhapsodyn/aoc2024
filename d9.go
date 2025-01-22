package main

import (
	"fmt"
	"strconv"
	"strings"
)

func d9p1(s string) int {
	s = strings.Trim(s, "\n")
	// -1 stands for free block
	var blocks []int
	dataCount := 0
	for i, r := range []rune(s) {
		n := Must(strconv.Atoi(string(r)))
		id := -1
		if i%2 == 0 {
			// data
			id = i / 2
			dataCount += n
		}

		for j := 0; j < n; j++ {
			blocks = append(blocks, id)
		}
	}
	// fmt.Println(blocks)

	// rearrange
	lastDataIdx := len(blocks) - 1
	for i := 0; i < dataCount; i++ {
		if blocks[i] == -1 {
			for ; blocks[lastDataIdx] == -1; lastDataIdx-- {
			}
			// swap
			blocks[i] = blocks[lastDataIdx]
			blocks[lastDataIdx] = -1
		}
	}
	// fmt.Println(blocks)

	// checksum
	sum := 0
	for i := 0; i < dataCount; i++ {
		sum += blocks[i] * i
	}

	return sum
}

type Block9 struct {
	id   int
	span int
}

func pretty9(blocks []Block9) {
	for _, b := range blocks {
		for i := 0; i < b.span; i++ {
			if b.id == -1 {
				fmt.Print(".")
			} else {
				fmt.Print(b.id)
			}
		}
	}
	fmt.Println()

}

func d9p2(s string) int {
  s = strings.Trim(s, "\n")

	var blocks []Block9
	var lastId int
	for i, r := range []rune(s) {
		n := Must(strconv.Atoi(string(r)))
		id := -1
		if i%2 == 0 {
			// data
			id = i / 2
			lastId = id
		}

		blocks = append(blocks, Block9{id: id, span: n})
	}
	// fmt.Println("orig:")
	//  pretty9(blocks)

	// rearrange
	// id == 0 has no effect on result
	for backId := lastId; backId > 0; backId-- {
		back := len(blocks) - 1
		for {
			if blocks[back].id == backId {
				break
			}
			back--
		}

		for front := 0; front < back; front++ {
			if blocks[front].id == -1 && blocks[front].span >= blocks[back].span {
				infill := blocks[front].span - blocks[back].span
				// fmt.Println("front:", front, "infill:", infill)
				// swap
				blocks[front].id = blocks[back].id
				blocks[front].span = blocks[back].span
				blocks[back].id = -1
				// may split
				if infill > 0 {
					var newBlocks []Block9
					newBlocks = append(newBlocks, blocks[:front+1]...)
					newBlocks = append(newBlocks, Block9{id: -1, span: infill})
					newBlocks = append(newBlocks, blocks[front+1:]...)
					blocks = newBlocks

					// pretty9(blocks)
					// panic("dbg")
				}

				break
			}
		}
	}
	// fmt.Println("after frag:")
	// pretty9(blocks)

	score := 0
	n := 0
	for _, b := range blocks {
		for i := 0; i < b.span; i++ {
			if b.id != -1 {
        score += n * b.id
			}
			n++
		}
	}

	return score
}
