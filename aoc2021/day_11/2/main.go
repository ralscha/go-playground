package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type octo struct {
	value   int32
	flashed bool
}

var grid [][]*octo

func main() {
	inFile, err := os.Open("./day_11/input.dat")
	if err != nil {
		log.Fatalf("loading input data failed: %v", err)
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() {
		line := scanner.Text()
		var row []*octo
		for _, c := range line {
			row = append(row, &octo{
				value:   c - '0',
				flashed: false,
			})
		}
		grid = append(grid, row)
	}

	round := 1
	for {
		for _, row := range grid {
			for _, col := range row {
				col.value++
			}
		}

		flashes := 0
		hashFlashes := true
		for hashFlashes {
			hashFlashes = false
			for rowIx, row := range grid {
				for colIx, col := range row {
					if col.value > 9 {
						col.value = 0
						flashes++
						incrNeighbors(rowIx, colIx)
						hashFlashes = true
					}
				}
			}
		}
		if flashes == 100 {
			fmt.Println("Round: ", round)
			break
		}
		round++
	}

}

func incrValue(rowIx, colIx int) {
	if grid[rowIx][colIx].value > 0 {
		grid[rowIx][colIx].value++
	}
}

func incrNeighbors(rowIx, colIx int) {
	if rowIx > 0 {
		if colIx > 0 {
			incrValue(rowIx-1, colIx-1)
		}
		incrValue(rowIx-1, colIx)

		if colIx < len(grid[0])-1 {
			incrValue(rowIx-1, colIx+1)
		}
	}

	if colIx > 0 {
		incrValue(rowIx, colIx-1)
	}
	if colIx < len(grid[0])-1 {
		incrValue(rowIx, colIx+1)
	}

	if rowIx < len(grid)-1 {
		if colIx > 0 {
			incrValue(rowIx+1, colIx-1)
		}
		incrValue(rowIx+1, colIx)
		if colIx < len(grid[0])-1 {
			incrValue(rowIx+1, colIx+1)
		}
	}
}
