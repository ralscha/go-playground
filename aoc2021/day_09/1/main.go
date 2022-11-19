package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	inFile, err := os.Open("./day_09/input.dat")
	if err != nil {
		log.Fatalf("loading input data failed: %v", err)
	}
	defer func(inFile *os.File) {
		err := inFile.Close()
		if err != nil {
			log.Fatalf("closing input file failed: %v", err)
		}
	}(inFile)

	var grid [][]int32

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		var lg []int32
		for _, c := range line {
			lg = append(lg, c-'0')
		}
		grid = append(grid, lg)
	}

	var lowPoints []int32
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			currentValue := grid[row][col]
			isLower := true
			//up
			if row > 0 {
				isLower = isLower && currentValue < grid[row-1][col]
			}
			//down
			if row < len(grid)-1 {
				isLower = isLower && currentValue < grid[row+1][col]
			}
			//left
			if col > 0 {
				isLower = isLower && currentValue < grid[row][col-1]
			}
			//right
			if col < len(grid[row])-1 {
				isLower = isLower && currentValue < grid[row][col+1]
			}

			if isLower {
				lowPoints = append(lowPoints, currentValue)
			}
		}
	}

	var total int32
	for _, lp := range lowPoints {
		total += lp + 1
	}

	fmt.Println("Result: ", total)
}
