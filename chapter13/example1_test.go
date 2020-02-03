package chapter13

import (
	"fmt"
	"testing"
)

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
