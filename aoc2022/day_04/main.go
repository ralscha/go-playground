package main

import (
	"adventofcode.com/2022/internal/conv"
	"adventofcode.com/2022/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	inputFile := "./day_04/input.txt"
	input, err := download.ReadInput(inputFile, 2022, 4)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	fullyContain := 0
	for _, line := range strings.Split(input, "\n") {
		splitted := strings.Split(line, ",")
		if len(splitted) != 2 {
			continue
		}
		firstPair := splitted[0]
		secondPair := splitted[1]
		firstPairSplitted := strings.Split(firstPair, "-")
		secondPairSplitted := strings.Split(secondPair, "-")
		firstPairFirst := conv.MustAtoi(firstPairSplitted[0])
		firstPairSecond := conv.MustAtoi(firstPairSplitted[1])
		secondPairFirst := conv.MustAtoi(secondPairSplitted[0])
		secondPairSecond := conv.MustAtoi(secondPairSplitted[1])

		if firstPairFirst <= secondPairFirst && firstPairSecond >= secondPairSecond {
			fullyContain += 1
		} else if secondPairFirst <= firstPairFirst && secondPairSecond >= firstPairSecond {
			fullyContain += 1
		}
	}
	fmt.Println(fullyContain)
}

func part2(input string) {
	partiallyContain := 0
	for _, line := range strings.Split(input, "\n") {
		splitted := strings.Split(line, ",")
		if len(splitted) != 2 {
			continue
		}
		firstPair := splitted[0]
		secondPair := splitted[1]
		firstPairSplitted := strings.Split(firstPair, "-")
		secondPairSplitted := strings.Split(secondPair, "-")
		firstPairFirst := conv.MustAtoi(firstPairSplitted[0])
		firstPairSecond := conv.MustAtoi(firstPairSplitted[1])
		secondPairFirst := conv.MustAtoi(secondPairSplitted[0])
		secondPairSecond := conv.MustAtoi(secondPairSplitted[1])

		if firstPairFirst <= secondPairFirst && firstPairSecond >= secondPairFirst {
			partiallyContain += 1
		} else if secondPairFirst <= firstPairFirst && secondPairSecond >= firstPairFirst {
			partiallyContain += 1
		}
	}
	fmt.Println(partiallyContain)
}
