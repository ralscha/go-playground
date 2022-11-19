package main

import "fmt"

func main() {
	bus := newGraph(true, true)

	vertices := make([]*vertex, 5)
	for i := 0; i < 5; i++ {
		vertices[i] = bus.addVertex(string(rune(i + 65)))
	}

	startVertex := vertices[0]

	bus.addEdge(vertices[0], vertices[1], 1)
	bus.addEdge(vertices[1], vertices[2], 4)
	bus.addEdge(vertices[1], vertices[3], 3)
	bus.addEdge(vertices[2], vertices[4], 6)

	bus.printGraph()
	bus.dfs(startVertex)
	fmt.Println()
	bus.bfs(startVertex)
	fmt.Println()
	distances := bus.dijkstra(startVertex)
	for v, d := range distances {
		fmt.Printf("distance from %s to %s is %d\n", startVertex.data, v.data, d)
	}

}
