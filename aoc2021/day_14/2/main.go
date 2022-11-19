package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {

	inFile, err := os.Open("./day_14/input.dat")
	if err != nil {
		log.Fatalf("loading input data failed: %v", err)
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	scanner.Scan()
	template := scanner.Text()
	rules := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		pos := strings.Index(line, " -> ")
		if pos != -1 {
			rules[line[:pos]] = line[pos+4:]
		}
	}

	pairs := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		pairs[template[i:i+2]] += 1
	}

	for i := 0; i < 40; i++ {
		newpairs := make(map[string]int)
		for k, count := range pairs {
			if v, ok := rules[k]; ok {
				newpairs[k[:1]+v] += count
				newpairs[v+k[1:]] += count
			}
		}
		pairs = newpairs
	}

	chars := make(map[string]int)
	for k, v := range pairs {
		chars[k[:1]] += v
		chars[k[1:]] += v
	}

	max, min := math.MinInt, math.MaxInt

	for _, v := range chars {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	fmt.Println("Result: ", (max-min)/2+1)
}
