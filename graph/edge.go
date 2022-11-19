package main

type edge struct {
	start  *vertex
	end    *vertex
	weight int
}

func newEdge(start, end *vertex, weight int) *edge {
	return &edge{start: start, end: end, weight: weight}
}
