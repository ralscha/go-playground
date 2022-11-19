package main

import "fmt"

type vertex struct {
	data  string
	edges []edge
}

func newVertex(data string) *vertex {
	return &vertex{data: data}
}

func (v *vertex) addEdge(end *vertex, weight int) {
	v.edges = append(v.edges, edge{start: v, end: end, weight: weight})
}

func (v *vertex) removeEdge(end *vertex) {
	for i, e := range v.edges {
		if e.end == end {
			v.edges = append(v.edges[:i], v.edges[i+1:]...)
			break
		}
	}
}

func (v *vertex) printEdges(printWeight bool) {
	if len(v.edges) == 0 {
		fmt.Printf("%s --> ", v.data)
		return
	}

	for _, e := range v.edges {
		if printWeight {
			fmt.Printf("%s --(%d)--> %s\n", e.start.data, e.weight, e.end.data)
		} else {
			fmt.Printf("%s --> %s\n", e.start.data, e.end.data)
		}
	}

}
