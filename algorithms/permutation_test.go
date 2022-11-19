package main

import "testing"

func BenchmarkPermutation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		permutation("abc")
	}
}

func BenchmarkPermutationIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		permutationIterative("abc")
	}
}
