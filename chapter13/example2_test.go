package chapter13

import (
	"fmt"
	"testing"
)

var gbs string

func BenchmarkPrint(b *testing.B) {
	b.Run("none", benchmarkSprint)
	b.Run("format", benchmarkSprintf)
}

func benchmarkSprint(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = fmt.Sprint("hello")
	}
	gbs = s
}

func benchmarkSprintf(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("hello")
	}
	gbs = s
}
