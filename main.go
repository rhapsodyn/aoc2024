package main

import (
	"fmt"
	"io"
	"os"
	"runtime/pprof"
)

func readInput(n int) string {
	name := fmt.Sprintf("inputs/%d.txt", n)
	bytes := Must(io.ReadAll(Must(os.Open(name))))
	return string(bytes)
}

func run() {
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
	// a := d5p2(readInput(5))
	// a := d6p1(readInput(6))
	// a := d6p2(readInput(6))
	// a := d7p1(readInput(7))
	// a := d7p2(readInput(7))
	// a := d8p1(readInput(8))
	// a := d8p2(readInput(8))
	// a := d9p1(readInput(9))
	// a := d9p2(readInput(9))
	// a := dap1(readInput(10))
	// a := dap2(readInput(10))
	// a := dbp1(readInput(11), 25)
	a := dbp2(readInput(11), 75)
	fmt.Println("The ANSWER is: ")
	fmt.Println(a)
}

func main() {
	args := os.Args
	if len(args) > 1 && args[1] == "prof" {
		f := Must(os.Create("cpu.prof"))
		defer f.Close()
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	run()
}
