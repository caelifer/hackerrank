// Go implementation for the "Functions and Fractals - Recursive Trees - Bash!" problem.
// https://www.hackerrank.com/challenges/fractal-trees-all/problem
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
	scn := NewScreen()
	for z := uint64(0); z < 6; z++ {
		x := uint64(1 << z)
		for i := uint64(0); i < h/x+1; i++ {
			scn.AppendLine(stem(x), x)
		}
		for i := uint64(0); i < h/x+1; i++ {
			scn.AppendLine(fork(x, i), x)
		}
	}
	scn.Print()
}

func stem(scale uint64) uint64 {
	return 1 << (64/(2*scale) - 1)
}

func fork(scale, iter uint64) uint64 {
	n := uint64((1 << (2 * iter)) | 1)
	return n << (64/(2*scale) - iter - 1)
}

type Row uint64

func (r Row) String() string {
	// Convert uint64 to []byte{} in-place, replacing 0 with '_' an 1 with '1'.
	buf := [64]byte{}
	for i := len(buf) - 1; i >= 0; i-- {
		buf[i] = '_' - byte((r>>i)&0x1*46) // magic :)
	}
	return string(buf[:])
}

type Screen []Row

func NewScreen() *Screen {
	s := Screen(make([]Row, 0, 64))
	return &s
}

func (s Screen) Print() {
	fmt.Println(topLine)
	// Print in reverse, skipping last line.
	for i := len(s) - 2; i > 0; i-- {
		fmt.Printf("%s%s%s\n", sidePad, s[i], sidePad)
	}
}

func (s *Screen) AppendLine(seg, scale uint64) {
	n := uint64(0)
	for i := uint64(0); i < scale; i++ {
		mag := i * 64 / scale // shift magnitude
		n |= seg << mag
	}
	*s = append(*s, Row(n))
}
