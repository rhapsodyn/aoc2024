package main

import (
	"fmt"
	"io"
	"os"
)

func readInput(n int) string {
	name := fmt.Sprintf("inputs/%d.txt", n)
	bytes := Must(io.ReadAll(Must(os.Open(name))))
	return string(bytes)
}

func main() {
	fmt.Println("Ladies and gentlemen, (may not block 4ever)")
	// a := d1p1(readInput(1))
	// a := d1p2(readInput(1))
	// a := d2p1(readInput(2))
	// a := d2p2(readInput(2))
	// a := d3p1(readInput(3))
	// a := d3p2(readInput(3))
	// a := d4p1(readInput(4))
	// a := d4p2(readInput(4))
	// a := d5p1(readInput(5))
	a := d5p2(readInput(5))
	fmt.Println("The ANSWER is: ")
	fmt.Println(a)
}
