package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	inFile, err := os.Open("./day_10/input.dat")
	if err != nil {
		log.Fatalf("loading input data failed: %v", err)
	}
	defer inFile.Close()
	openP := "([{<"
	closeP := ")]}>"

	var scores []int
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()
		var stack []string
		invalid := false
		for _, c := range line {
			s := string(c)
			pos := strings.Index(openP, s)
			if pos != -1 {
				cl := string(closeP[pos])
				stack = append(stack, cl)
			} else {
				n := len(stack) - 1
				cl := stack[n]
				if s != cl {
					invalid = true
					break
				}
				stack = stack[:n]
			}
		}
		if !invalid && len(stack) > 0 {
			score := 0
			for len(stack) > 0 {
				score = score * 5
				n := len(stack) - 1
				cl := stack[n]
				switch cl {
				case ")":
					score += 1
				case "]":
					score += 2
				case "}":
					score += 3
				case ">":
					score += 4
				}
				stack = stack[:n]
			}

			scores = append(scores, score)
		}
	}

	sort.Ints(scores)
	fmt.Println("Result: ", scores[len(scores)/2])

}
