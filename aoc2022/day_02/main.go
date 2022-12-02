package main

import (
	"adventofcode.com/2022/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	inputFile := "./day_02/input.txt"
	input, err := download.ReadInput(inputFile, 2022, 2)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	//9177
	//12111
	part1(input)
	part2(input)
}

func part1(input string) {
	totalScore := 0
	for _, line := range strings.Split(input, "\n") {
		in := strings.Split(line, " ")
		if len(in) == 2 {
			totalScore += score(in[0], in[1])
		}
	}
	fmt.Println(totalScore)
}

func part2(input string) {
	// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win
	totalScore := 0
	for _, line := range strings.Split(input, "\n") {
		in := strings.Split(line, " ")
		if len(in) == 2 {
			opponent := in[0]
			expected := in[1]
			myMove := ""
			if expected == "Y" {
				if opponent == "A" {
					myMove = "X"
				} else if opponent == "B" {
					myMove = "Y"
				} else {
					myMove = "Z"
				}
			} else if expected == "X" {
				if opponent == "A" {
					myMove = "Z"
				} else if opponent == "B" {
					myMove = "X"
				} else {
					myMove = "Y"
				}
			} else if expected == "Z" {
				if opponent == "A" {
					myMove = "Y"
				} else if opponent == "B" {
					myMove = "Z"
				} else {
					myMove = "X"
				}
			}
			totalScore += score(opponent, myMove)
		}
	}
	fmt.Println(totalScore)
}

func score(o, m string) int {
	// 1 for Rock, 2 for Paper, and 3 for Scissors
	score := 0
	if m == "X" {
		score += 1
	} else if m == "Y" {
		score += 2
	} else if m == "Z" {
		score += 3
	}

	// 0 if you lost, 3 if the round was a draw, and 6 if you won
	if o == "A" {
		if m == "X" {
			score += 3
		} else if m == "Y" {
			score += 6
		} else if m == "Z" {
			score += 0
		}
	} else if o == "B" {
		if m == "X" {
			score += 0
		} else if m == "Y" {
			score += 3
		} else if m == "Z" {
			score += 6
		}
	} else if o == "C" {
		if m == "X" {
			score += 6
		} else if m == "Y" {
			score += 0
		} else if m == "Z" {
			score += 3
		}
	}

	return score
}
