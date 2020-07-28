// Go implementation for the "Functions and Fractals - Recursive Trees - Bash!" problem.
// https://www.hackerrank.com/challenges/fractal-trees-all/problem
package main

import (
	"fmt"
	"strings"
)

const h = 15 // half-heigth of the Y pattern

func main() {
	var scrn = []string{}
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

func addLine(slice []string, seg, scale uint64) []string {
	n, ln := uint64(0), 64/scale

	for i := uint64(0); i < scale; i++ {
		n |= seg << (i * ln)
	}
	line := strings.ReplaceAll(fmt.Sprintf("%064b", n), "0", "_")
	return append(slice, line)
}

func printScreen(scrn []string) {
	fmt.Println("____________________________________________________________________________________________________")
	for i := len(scrn) - 2; i > 0; i-- {
		fmt.Printf("%s%s%s\n", "_________________", scrn[i], "___________________")
	}
}
