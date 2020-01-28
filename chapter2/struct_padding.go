package main

import (
	"fmt"
	"unsafe"
)

// This causes a padding of extra 21 bytes.
// However this is desirable since this improves the readability of the struct.
type foo struct {
	flag1 bool
	count1 float64
	flag2 bool
	count2 float64
	flag3 bool
	count3 float64
}

// This is pre-optimization of the fields on the struct.
// Although the size reduces, this hampers the productivity.
type compactFoo struct {
	count1 float64
	count2 float64
	count3 float64
	flag1 bool
	flag2 bool
	flag3 bool
}

func main() {
var f1 foo
fmt.Printf("foo size %d", unsafe.Sizeof(f1))
var f2 compactFoo
fmt.Printf("compactfoo size %d", unsafe.Sizeof(f2))
}
