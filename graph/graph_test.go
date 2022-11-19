package main

import "testing"

func TestNewGraph(t *testing.T) {
	// act
	g := newGraph(true, false)

	// assert
	if g.isWeighted != true {
		t.Errorf("Expected isWeighted to be %v, but got %v", true, g.isWeighted)
	}
	if g.isDirected != false {
		t.Errorf("Expected isDirected to be %v, but got %v", false, g.isDirected)
	}
}

func TestGraphAddVertex(t *testing.T) {
	// arrange
	g := newGraph(true, false)

	// act
	v := g.addVertex("Clifton")

	// assert
	if len(g.vertices) != 1 {
		t.Errorf("Expected 1 vertex, but got %v", len(g.vertices))
	}

	if v.data != "Clifton" {
		t.Errorf("Expected data to be Clifton, but got %v", v.data)
	}
}

func TestGraphAddEdgeNotDirected(t *testing.T) {
	// arrange
	g := newGraph(true, false)
	start := g.addVertex("Clifton")
	end := g.addVertex("Cape May")
	weight := 1000

	// act
	g.addEdge(start, end, weight)

	// assert
	if len(start.edges) != 1 {
		t.Errorf("Expected 1 edge, but got %v", len(start.edges))
	}
	if len(end.edges) != 1 {
		t.Errorf("Expected 0 edges, but got %v", len(end.edges))
	}
}

func TestGraphAddEdgeDirected(t *testing.T) {
	// arrange
	g := newGraph(true, true)
	start := g.addVertex("Clifton")
	end := g.addVertex("Cape May")
	weight := 1000

	// act
	g.addEdge(start, end, weight)

	// assert
	if len(start.edges) != 1 {
		t.Errorf("Expected 1 edge, but got %v", len(start.edges))
	}
	if len(end.edges) != 0 {
		t.Errorf("Expected 1 edge, but got %v", len(end.edges))
	}
}

func TestGraphRemoveEdgeNotDirected(t *testing.T) {
	// arrange
	g := newGraph(true, false)
	start := g.addVertex("Clifton")
	end := g.addVertex("Cape May")
	g.addEdge(start, end, 1000)

	// act
	g.removeEdge(start, end)

	// assert
	if len(start.edges) != 0 {
		t.Errorf("Expected 0 edges, but got %v", len(start.edges))
	}
	if len(end.edges) != 0 {
		t.Errorf("Expected 0 edges, but got %v", len(end.edges))
	}
}

func TestGraphRemoveEdgeDirected(t *testing.T) {
	// arrange
	g := newGraph(true, true)
	start := g.addVertex("Clifton")
	end := g.addVertex("Cape May")
	g.addEdge(start, end, 1000)

	// act
	g.removeEdge(start, end)

	// assert
	if len(start.edges) != 0 {
		t.Errorf("Expected 0 edges, but got %v", len(start.edges))
	}
	if len(end.edges) != 0 {
		t.Errorf("Expected 0 edges, but got %v", len(end.edges))
	}
}

func TestGraphRemoveVertex(t *testing.T) {
	// arrange
	g := newGraph(true, false)
	v := g.addVertex("Clifton")

	// act
	g.removeVertex(v)

	// assert
	if len(g.vertices) != 0 {
		t.Errorf("Expected 0 vertices, but got %v", len(g.vertices))
	}
}

func TestGraphDijkstra(t *testing.T) {
	// arrange
	g := newGraph(true, true)

	vertices := make([]*vertex, 5)
	for i := 0; i < 5; i++ {
		vertices[i] = g.addVertex(string(rune(i + 65)))
	}
	start := vertices[0]
	g.addEdge(vertices[0], vertices[1], 1)
	g.addEdge(vertices[1], vertices[2], 4)
	g.addEdge(vertices[1], vertices[3], 3)
	g.addEdge(vertices[2], vertices[4], 6)

	// act
	distances := g.dijkstra(start)

	// assert
	if distances[vertices[0]] != 0 {
		t.Errorf("Expected distance to be 0, but got %v", distances[vertices[0]])
	}
	if distances[vertices[1]] != 1 {
		t.Errorf("Expected distance to be 1, but got %v", distances[vertices[1]])
	}
	if distances[vertices[2]] != 5 {
		t.Errorf("Expected distance to be 5, but got %v", distances[vertices[2]])
	}
	if distances[vertices[3]] != 4 {
		t.Errorf("Expected distance to be 4, but got %v", distances[vertices[3]])
	}
	if distances[vertices[4]] != 11 {
		t.Errorf("Expected distance to be 11, but got %v", distances[vertices[4]])
	}
}
