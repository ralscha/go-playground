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

	inFile, err := os.Open("./day_05/input.dat")
	if err != nil {
		log.Fatalf("loading input data failed: %v", err)
	}
	defer func(inFile *os.File) {
		err := inFile.Close()
		if err != nil {
			log.Fatalf("closing input file failed: %v", err)
		}
	}(inFile)

	var grid [1000][1000]int
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()

		splitted := strings.Split(line, " -> ")

		startSplitted := strings.Split(splitted[0], ",")
		endSplitted := strings.Split(splitted[1], ",")
		startX, err := strconv.Atoi(startSplitted[0])
		if err != nil {
			log.Fatalf("conversion failed: %s %v", startSplitted[0], err)
		}
		startY, err := strconv.Atoi(startSplitted[1])
		if err != nil {
			log.Fatalf("conversion failed: %s %v", startSplitted[1], err)
		}
		endX, err := strconv.Atoi(endSplitted[0])
		if err != nil {
			log.Fatalf("conversion failed: %s %v", endSplitted[0], err)
		}
		endY, err := strconv.Atoi(endSplitted[1])
		if err != nil {
			log.Fatalf("conversion failed: %s %v", endSplitted[1], err)
		}

		if startX == endX {
			start, end := startY, endY
			if start > end {
				start, end = end, start
			}
			for y := start; y <= end; y++ {
				grid[y][startX] += 1
			}
		} else if startY == endY {
			start, end := startX, endX
			if start > end {
				start, end = end, start
			}
			for x := start; x <= end; x++ {
				grid[startY][x] += 1
			}
		}
	}

	count := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] >= 2 {
				count++
			}
		}
	}

	fmt.Println("Result: ", count)
}
