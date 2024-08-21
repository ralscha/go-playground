package main

import (
	"fmt"
	"testing"
)

func BenchmarkHistory(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("Add_Size_%d", size), func(b *testing.B) {
			h := NewHistory(size)
			b.ResetTimer()
			for i := 0; i < 10*size*b.N; i++ {
				h.Add("benchmark")
				_ = h.String()
			}
		})
	}
}
