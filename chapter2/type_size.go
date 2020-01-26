package main

import (
	"fmt"
	"unsafe"
)
// Total size be 16 bytes; 1 byte padding after bool and 4 bytes after int16.
// The size of struct is 16.
type foo1 struct {
	bdummy bool
	idummy int16 // padding for this would only be 1 byte.
	dummy int
}

// The struct will align on boundary of size of biggest type;
// The size of struct will be 4.
type foo2 struct {
	float int16
	bool bool
}

// The struct will align on boundary of size of biggest type;
// The size of struct will be 2.
type foo3 struct {
	bool1 bool
	bool2 bool
}

func main() {
	f := foo1{}
	fmt.Printf("Size of foo is %d",unsafe.Sizeof(f))
	f2 := foo2{}
	fmt.Printf("Size of foo is %d",unsafe.Sizeof(f2))
	f3 := foo3{}
	fmt.Printf("Size of foo is %d",unsafe.Sizeof(f3))
}
