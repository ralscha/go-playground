package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	inFile, err := os.Open("./day_07/input.dat")
	if err != nil {
		log.Fatalf("loading input data failed: %v", err)
	}
	defer func(inFile *os.File) {
		err := inFile.Close()
		if err != nil {
			log.Fatalf("closing input file failed: %v", err)
		}
	}(inFile)

	scanner := bufio.NewScanner(inFile)
	scanner.Scan()
	line := scanner.Text()
	splitted := strings.Split(line, ",")

	var crabs []int
	for _, s := range splitted {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("conversion failed: %s %v", s, err)
		}

		crabs = append(crabs, n)
	}

	leastFuel := -1
	leastPos := -1

	min, max := findMinMax(crabs)
	for p := min; p <= max; p++ {
		fuel := 0
		for _, c := range crabs {
			fuel += abs(c - p)
		}
		if fuel < leastFuel || leastFuel == -1 {
			leastFuel = fuel
			leastPos = p
		}
	}

	fmt.Println("Result: ", leastPos, leastFuel)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findMinMax(a []int) (int, int) {
	min := a[0]
	max := a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}
