package main

import (
	"adventofcode.com/2022/internal/download"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputFile := "./day_01/input.txt"
	input, err := download.ReadInput(inputFile, 2022, 1)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	var elves []int
	totalCalories := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			elves = append(elves, totalCalories)
			totalCalories = 0
		} else {
			calories, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalf("converting input to int failed: %v", err)
			}
			totalCalories += calories
		}
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})

	fmt.Printf("Largest amount of calories is %d\n", elves[0])

	sum := 0
	for i := 0; i < 3; i++ {
		sum += elves[i]
	}
	fmt.Printf("Calories of top 3 elves is %d\n", sum)
}
