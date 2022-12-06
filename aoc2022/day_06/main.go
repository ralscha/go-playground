package main

import (
	"adventofcode.com/2022/internal/download"
	"log"
)

func main() {
	inputFile := "./day_06/input.txt"
	input, err := download.ReadInput(inputFile, 2022, 6)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input, 4)
	part1and2(input, 14)
}

func part1and2(input string, blockLength int) {
	if len(input) < blockLength {
		return
	}

	for i := 0; i < len(input)-blockLength; i++ {
		block := input[i : i+blockLength]
		if uniqueCharacters(block) {
			log.Printf("index: %d", i+blockLength)
			return
		}
	}
}

func uniqueCharacters(s string) bool {
	// return true if string contains unique characters
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}
