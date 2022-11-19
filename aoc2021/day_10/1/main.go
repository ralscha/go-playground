package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	var stack []string

	scanner := bufio.NewScanner(inFile)
	score := 0
	for scanner.Scan() {
		line := scanner.Text()
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
					log.Printf("Expected %s, but found %s instead.", cl, s)
					switch s {
					case ")":
						score += 3
					case "]":
						score += 57
					case "}":
						score += 1197
					case ">":
						score += 25137
					}
					break
				}
				stack = stack[:n]
			}
		}
	}

	fmt.Println("Result: ", score)

}
