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

	inFile, err := os.Open("./day_02/input.dat")
	if err != nil {
		log.Fatalf("loading input data failed: %v", err)
	}
	defer inFile.Close()

	horizontalPosition := 0
	depth := 0

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Fields(line)
		num, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatalf("conversion failed: %v", err)
		}
		switch split[0] {
		case "down":
			depth += num
		case "up":
			depth -= num
		case "forward":
			horizontalPosition += num
		}
	}

	fmt.Println("Depth: ", depth)
	fmt.Println("Horizontal Position: ", horizontalPosition)
	fmt.Println("Result: ", depth*horizontalPosition)

}
