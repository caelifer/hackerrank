// Go implementation for the "Functions and Fractals - Recursive Trees - Bash!" problem.
// https://www.hackerrank.com/challenges/fractal-trees-all/problem
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	h       = 15 // half-heigth of the Y pattern
	sidePad = "_________________"
	topLine = "__________________________________________________________________________________________________"
)

func main() {
	if err := run(os.Stdout); err != nil {
		log.Fatal(err)
	}
}

func run(w io.Writer) error {
	scn := NewScreen()
	for z := uint64(0); z < 6; z++ {
		x := uint64(1 << z)
		for i := uint64(0); i < h/x+1; i++ {
			scn.AppendLine(drawStemPart(x), x)
		}
		for i := uint64(0); i < h/x+1; i++ {
			scn.AppendLine(drawForkPart(x, i), x)
		}
	}
	_, err := fmt.Fprintln(w, scn)
	return err
}

func drawStemPart(scale uint64) uint64 {
	return 1 << (64/(2*scale) - 1)
}

func drawForkPart(scale, iter uint64) uint64 {
	n := uint64((1 << (2 * iter)) | 1)
	return n << (64/(2*scale) - iter - 1)
}

type Row uint64

func (r Row) String() string {
	// Convert uint64 to []byte{} in-place by replacing 0 with '_' an 1 with '1' ASCII code representations.
	buf := [64]byte{}
	for i := len(buf) - 1; i >= 0; i-- {
		buf[i] = '_' - byte((r>>i)&0x1)*46 // magic :)
	}
	return string(buf[:])
}

type Screen []Row

func NewScreen() *Screen {
	s := Screen(make([]Row, 0, 64))
	return &s
}

func (s Screen) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprintln(buf, topLine)
	// Print in reverse, skipping last line.
	for i := len(s) - 2; i > 0; i-- {
		fmt.Fprintf(buf, "%s%s%s\n", sidePad, s[i], sidePad)
	}
	return buf.String()
}

func (s *Screen) AppendLine(seg, scale uint64) {
	n := uint64(0)
	for i := uint64(0); i < scale; i++ {
		mag := i * 64 / scale // shift magnitude
		n |= seg << mag
	}
	*s = append(*s, Row(n))
}

// vim: :ts=4:sw=4:noet:ai
