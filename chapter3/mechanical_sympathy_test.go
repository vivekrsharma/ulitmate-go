package main

import "testing"

var fa int
func BenchmarkLinkedListTraversal( b *testing.B) {
	var a int
	for i:=0;i<b.N;i++ {
		a = linkedListTraversal();
	}
	fa = a
}

func BenchmarkRowWiseTraversal(b *testing.B) {
	var a int
	for i:=0;i <b.N;i++ {
		a = rowWiseTraversal()
		fa = a
	}
}

func BenchmarkColumnWiseTraversal(b *testing.B) {
	var a int
	for i:=0;i <b.N;i++ {
		a = columnWiseTraversal()
		fa = a
	}
}