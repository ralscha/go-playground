package main

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkStringBuilderConcatenation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		for range 1000 {
			sb.WriteString("h")
		}
		_ = sb.String()
	}
}

func BenchmarkStringConcatenation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s string
		for range 1000 {
			s += "h"
		}
	}
}

func BenchmarkFmtSprintfConcatenation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s string
		for range 1000 {
			s = fmt.Sprintf("%s%s", s, "h")
		}
	}
}
