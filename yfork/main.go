// Go implementation for the "Functions and Fractals - Recursive Trees - Bash!" problem.
// https://www.hackerrank.com/challenges/fractal-trees-all/problem
//
// Live code - https://play.golang.org/p/qXfP_6i7AMW
//
package main

import (
	"fmt"
)

const (
	h       = 15 // half-heigth of the Y pattern
	sidePad = "_________________"
	topLine = "____________________________________________________________________________________________________"
)

func main() {
	var scrn = [][64]byte{}
	for z := uint64(0); z < 6; z++ {
		x := uint64(1 << z)
		for i := uint64(0); i < h/x+1; i++ {
			scrn = addLine(scrn, stem(x), x)
		}
		for i := uint64(0); i < h/x+1; i++ {
			scrn = addLine(scrn, fork(x, i), x)
		}
	}
	printScreen(scrn)
}

func stem(scale uint64) uint64 {
	return 1 << (64/(2*scale) - 1)
}

func fork(scale, iter uint64) uint64 {
	n := uint64((1 << (2 * iter)) | 1)
	return n << (64/(2*scale) - iter - 1)
}

func addLine(scrBuf [][64]byte, seg, scale uint64) [][64]byte {
	n, ln := uint64(0), 64/scale

	for i := uint64(0); i < scale; i++ {
		n |= seg << (i * ln)
	}

	// Convert uint64 to []byte{} in-place, replacing 0 with '_' an 1 with '1'.
	buf := [64]byte{}
	for i := len(buf) - 1; i >= 0; i-- {
		if (n>>i)&0x1 != 0 {
			buf[i] = '1'
		} else {
			buf[i] = '_'
		}
	}

	// Add a line to a screen buffer.
	return append(scrBuf, buf)
}

func printScreen(scrn [][64]byte) {
	fmt.Println(topLine)
	for i := len(scrn) - 2; i > 0; i-- {
		fmt.Printf("%s%s%s\n", sidePad, string(scrn[i][:]), sidePad)
	}
}
