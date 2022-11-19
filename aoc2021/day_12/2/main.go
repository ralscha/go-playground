package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

type node struct {
	start      bool
	end        bool
	name       string
	visitCount int
}

var nodes = make(map[string]*node)
var adjmap = make(map[*node][]*node)

func hasTwo() *node {
	for _, n := range nodes {
		if !isUpper(n.name) && !n.end && !n.start && n.visitCount >= 2 {
			return n
		}
	}
	return nil
}

func (n *node) canVisit() bool {
	if n.end || n.start {
		return true
	}
	if isUpper(n.name) {
		return true
	}
	if n.visitCount == 0 {
		return true
	}
	ht := hasTwo()
	if ht == nil {
		return true
	}
	if n.visitCount == 1 && ht == n {
		return true
	}
	return false
}

func addEdge(u, v *node) {
	if _, ok := adjmap[u]; ok {
		adjmap[u] = append(adjmap[u], v)
	} else {
		adjmap[u] = []*node{v}
	}
}

func isUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func countAllPaths(s, d *node) int {
	var pathList []string

	pathList = append(pathList, s.name)
	return printAllPathsUtil(s, d, pathList, 0)
}

func printAllPathsUtil(u, d *node, localPathList []string, count int) int {
	if u == d {
		return 1
	}

	u.visitCount++

	for _, i := range adjmap[u] {
		if i.canVisit() {
			localPathList = append(localPathList, i.name)
			ix := len(localPathList) - 1
			count += printAllPathsUtil(i, d, localPathList, 0)
			localPathList = append(localPathList[:ix], localPathList[ix+1:]...)
		}
	}

	u.visitCount--

	return count
}

func main() {

	inFile, err := os.Open("./day_12/input.dat")
	if err != nil {
		log.Fatalf("loading input data failed: %v", err)
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() {
		line := scanner.Text()
		ix := strings.Index(line, "-")
		start, end := line[:ix], line[ix+1:]

		startNode, ok := nodes[start]
		if !ok {
			startNode = &node{
				start:      start == "start",
				end:        start == "end",
				name:       start,
				visitCount: 0,
			}
			nodes[start] = startNode
		}

		endNode, ok := nodes[end]
		if !ok {
			endNode = &node{
				start:      end == "start",
				end:        end == "end",
				name:       end,
				visitCount: 0,
			}
			nodes[end] = endNode
		}
		if startNode.start || endNode.end {
			addEdge(startNode, endNode)
		} else if endNode.start || startNode.end {
			addEdge(endNode, startNode)
		} else {
			addEdge(startNode, endNode)
			addEdge(endNode, startNode)
		}
	}

	var sn *node
	var en *node

	for _, n := range nodes {
		if n.end {
			en = n
		} else if n.start {
			sn = n
		}
	}

	count := countAllPaths(sn, en)
	fmt.Println("Result: ", count)
}
