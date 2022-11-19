package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	inFile, err := os.Open("./day_16/input.dat")
	if err != nil {
		log.Fatalf("loading input data failed: %v", err)
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)

	grid := make(map[point]struct{})

	for scanner.Scan() {
		line := scanner.Text()
	}

	fmt.Println("Result: ", len(grid))
}
