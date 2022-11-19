package main

import "fmt"

type graph struct {
	vertices   []*vertex
	isWeighted bool
	isDirected bool
}

func newGraph(isWeighted, isDirected bool) *graph {
	return &graph{isWeighted: isWeighted, isDirected: isDirected}
}

func (g *graph) dfs(start *vertex) {
	visited := make(map[*vertex]struct{})
	g.dfsHelper(start, visited)
}

func (g *graph) dfsHelper(start *vertex, visited map[*vertex]struct{}) {
	visited[start] = struct{}{}
	fmt.Printf("%s ", start.data)
	for _, edge := range start.edges {
		if _, ok := visited[edge.end]; !ok {
			g.dfsHelper(edge.end, visited)
		}
	}
}

func (g *graph) bfs(start *vertex) {
	visited := make(map[*vertex]struct{})
	queue := []*vertex{start}
	visited[start] = struct{}{}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		fmt.Printf("%s ", current.data)
		for _, edge := range current.edges {
			if _, ok := visited[edge.end]; !ok {
				queue = append(queue, edge.end)
				visited[edge.end] = struct{}{}
			}
		}
	}
}

func (g *graph) dijkstra(start *vertex) map[*vertex]int {
	distances := make(map[*vertex]int)
	previous := make(map[*vertex]*vertex)
	visited := make(map[*vertex]struct{})
	unvisited := make(map[*vertex]struct{})

	for _, v := range g.vertices {
		distances[v] = int(^uint(0) >> 1)
	}

	distances[start] = 0

	for _, v := range g.vertices {
		unvisited[v] = struct{}{}
	}

	for len(unvisited) > 0 {
		var smallest *vertex
		for v := range unvisited {
			if smallest == nil {
				smallest = v
			} else if distances[v] < distances[smallest] {
				smallest = v
			}
		}

		visited[smallest] = struct{}{}
		delete(unvisited, smallest)
		for _, edge := range smallest.edges {
			alt := distances[smallest] + edge.weight
			if alt < distances[edge.end] {
				distances[edge.end] = alt
				previous[edge.end] = smallest
			}
		}
	}
	return distances
}

func (g *graph) addVertex(data string) *vertex {
	v := newVertex(data)
	g.vertices = append(g.vertices, v)
	return v
}

func (g *graph) addEdge(start, end *vertex, weight int) {
	w := weight
	if !g.isWeighted {
		w = 0
	}
	start.addEdge(end, w)
	if !g.isDirected {
		end.addEdge(start, w)
	}
}

func (g *graph) removeEdge(start, end *vertex) {
	start.removeEdge(end)
	if !g.isDirected {
		end.removeEdge(start)
	}
}

func (g *graph) removeVertex(v *vertex) {
	for i, vertex := range g.vertices {
		if vertex == v {
			g.vertices = append(g.vertices[:i], g.vertices[i+1:]...)
			break
		}
	}

	for _, vertex := range g.vertices {
		vertex.removeEdge(v)
	}
}

func (g *graph) getVertex(data string) *vertex {
	for _, v := range g.vertices {
		if v.data == data {
			return v
		}
	}
	return nil
}

func (g *graph) printGraph() {
	for _, v := range g.vertices {
		v.printEdges(g.isWeighted)
		fmt.Println()
	}
}
