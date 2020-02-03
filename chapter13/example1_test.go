package chapter13

import (
	"fmt"
	"testing"
)


// go test -run none -bench . -benchtime 3s -benchmem
// Dont guess - optimize for readability, simplicity, correctness, integrity. Analyse performance only after you're done.
var gb string
func BenchmarkSprint(b *testing.B) {
	var s string
	for i:=0;i<b.N;i++ {
		 s = fmt.Sprint("hello")
	}
	gb = s
}

func BenchmarkSprintf(b *testing.B) {
	var s string
	for i:=0;i<b.N;i++ {
		s = fmt.Sprintf("hello")
	}
	gb = s
}
