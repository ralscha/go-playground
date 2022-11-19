package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

var adjmap = make(map[string][]string)

func addEdge(u, v string) {
	if _, ok := adjmap[u]; ok {
		adjmap[u] = append(adjmap[u], v)
	} else {
		adjmap[u] = []string{v}
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

func countAllPaths(s, d string) int {
	visited := make(map[string]struct{})
	var pathList []string

	pathList = append(pathList, s)
	return printAllPathsUtil(s, d, visited, pathList, 0)
}

func printAllPathsUtil(u, d string, visited map[string]struct{}, localPathList []string, count int) int {
	if u == d {
		return 1
	}

	if !isUpper(u) {
		visited[u] = struct{}{}
	}

	for _, i := range adjmap[u] {
		if _, ok := visited[i]; !ok {
			localPathList = append(localPathList, i)
			ix := len(localPathList) - 1
			count += printAllPathsUtil(i, d, visited, localPathList, 0)
			localPathList = append(localPathList[:ix], localPathList[ix+1:]...)
		}
	}

	if !isUpper(u) {
		delete(visited, u)
	}

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
		if end == "end" || start == "start" {
			addEdge(start, end)
		} else {
			addEdge(start, end)
			addEdge(end, start)
		}
	}

	s, d := "start", "end"
	count := countAllPaths(s, d)
	fmt.Println("Result: ", count)
}
