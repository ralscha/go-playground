package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type field struct {
	number int
	marked bool
}

type card [5][5]*field

func main() {

	inFile, err := os.Open("./day_04/input.dat")
	if err != nil {
		log.Fatalf("loading input data failed: %v", err)
	}
	defer func(inFile *os.File) {
		err := inFile.Close()
		if err != nil {
			log.Fatalf("closing input file failed: %v", err)
		}
	}(inFile)

	// var input []string
	scanner := bufio.NewScanner(inFile)
	scanner.Scan()
	drawLine := scanner.Text()
	numberStrings := strings.Split(drawLine, ",")
	draw := make([]int, len(numberStrings))

	for ix, ns := range numberStrings {
		draw[ix], err = strconv.Atoi(ns)
		if err != nil {
			log.Fatalf("conversion from string to int failed: %s %v", ns, err)
		}
	}

	scanner.Scan()

	row := 0
	var cards []card
	var currentCard card

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			cards = append(cards, currentCard)
			row = 0
		} else {
			rowString := strings.Fields(line)

			for ix, ns := range rowString {
				n, err := strconv.Atoi(ns)
				if err != nil {
					log.Fatalf("conversion from string to int failed: %s %v", ns, err)
				}
				currentCard[row][ix] = &field{
					number: n,
					marked: false,
				}
			}

			row++
		}

	}

	var lastCalledNumber int
	var bingoCards []int

	for _, d := range draw {
		// mark
		for cardIx, card := range cards {
			if contains(bingoCards, cardIx) {
				continue
			}
			for _, row := range card {
				for _, col := range row {
					if col.number == d {
						col.marked = true
					}
				}
			}
		}

		// check wining
		bcs, ixs := checkBingo(cards, bingoCards)
		if len(bcs) > 0 {
			bingoCards = append(bingoCards, ixs...)
			lastCalledNumber = d
		}
	}
	fmt.Println(bingoCards)
	lastBingoCard := cards[bingoCards[len(bingoCards)-1]]
	fmt.Println(lastBingoCard)
	fmt.Println(lastCalledNumber)

	sumUnmarked := 0
	for _, row := range lastBingoCard {
		for _, col := range row {
			if !col.marked {
				sumUnmarked += col.number
			}
		}
	}

	fmt.Println("Result: ", sumUnmarked*lastCalledNumber)
}

func checkBingo(cards []card, bingoCards []int) ([]*card, []int) {
	var bcards []*card
	var bix []int

	for cardIx, card := range cards {
		if contains(bingoCards, cardIx) {
			continue
		}
		for _, row := range card {
			bingo := true
			for _, col := range row {
				if !col.marked {
					bingo = false
					break
				}
			}
			if bingo {
				bcards = append(bcards, &card)
				bix = append(bix, cardIx)
			}
		}

		for col := 0; col < 5; col++ {
			bingo := true
			for row := 0; row < 5; row++ {
				if !card[row][col].marked {
					bingo = false
					break
				}
			}
			if bingo && !contains(bix, cardIx) {
				bcards = append(bcards, &card)
				bix = append(bix, cardIx)
			}

		}
	}

	return bcards, bix
}

func contains(array []int, value int) bool {
	for _, b := range array {
		if b == value {
			return true
		}
	}
	return false
}
