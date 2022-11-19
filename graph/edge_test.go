package main

import (
	"testing"
)

func TestEdgeNewEdge(t *testing.T) {
	// arrange
	start := newVertex("Clifton")
	end := newVertex("Cape May")
	weight := 1000

	// act
	e := newEdge(start, end, weight)

	// assert
	if e.start != start {
		t.Errorf("Expected start to be %v, but got %v", start, e.start)
	}
	if e.end != end {
		t.Errorf("Expected end to be %v, but got %v", end, e.end)
	}
	if e.weight != weight {
		t.Errorf("Expected weight to be %v, but got %v", weight, e.weight)
	}
}
