package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	inFile, err := os.Open("./day_03/input.dat")
	if err != nil {
		log.Fatalf("loading input data failed: %v", err)
	}
	defer inFile.Close()

	var ones [12]int32
	var zeros [12]int32

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()
		for ix, c := range line {
			if c-'0' == 0 {
				zeros[ix] += 1
			} else {
				ones[ix] += 1
			}
		}
	}

	gamma := ""
	epsilon := ""

	for ix := range ones {
		if ones[ix] > zeros[ix] {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	gammaNumber, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		log.Fatalf("ParseInt failed: %s %v", gamma, err)
	}
	epsilonNumber, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		log.Fatalf("ParseInt failed: %s %v", epsilon, err)
	}
	fmt.Println("Result: ", gammaNumber*epsilonNumber)
}
