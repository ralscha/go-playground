package main

import "testing"

func TestVertex(t *testing.T) {
	// arrange
	v := newVertex("Clifton")

	// act
	v.addEdge(newVertex("Cape May"), 1000)

	// assert
	if len(v.edges) != 1 {
		t.Errorf("Expected 1 edge, but got %v", len(v.edges))
	}

	if v.data != "Clifton" {
		t.Errorf("Expected data to be Clifton, but got %v", v.data)
	}
}

func TestVertexAddEdge(t *testing.T) {
	// arrange
	v := newVertex("Clifton")
	end := newVertex("Cape May")
	weight := 1000

	// act
	v.addEdge(end, weight)

	// assert
	if len(v.edges) != 1 {
		t.Errorf("Expected 1 edge, but got %v", len(v.edges))
	}

	if v.edges[0].start != v {
		t.Errorf("Expected start to be %v, but got %v", v, v.edges[0].start)
	}

	if v.edges[0].end != end {
		t.Errorf("Expected end to be %v, but got %v", end, v.edges[0].end)
	}

	if v.edges[0].weight != weight {
		t.Errorf("Expected weight to be %v, but got %v", weight, v.edges[0].weight)
	}
}

func TestVertexRemoveEdge(t *testing.T) {
	// arrange
	v := newVertex("Clifton")
	end := newVertex("Cape May")
	v.addEdge(end, 1000)

	// act
	v.removeEdge(end)

	// assert
	if len(v.edges) != 0 {
		t.Errorf("Expected 0 edges, but got %v", len(v.edges))
	}
}
