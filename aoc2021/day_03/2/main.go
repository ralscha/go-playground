package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func count(lines []string, pos int, one bool) int {
	var ones int
	var zeros int
	for _, line := range lines {
		chars := []rune(line)
		c := chars[pos]
		if c-'0' == 0 {
			zeros += 1
		} else {
			ones += 1
		}
	}
	if one {
		if zeros > ones {
			return 0
		}
		return 1
	} else {
		if ones < zeros {
			return 1
		}
		return 0
	}
}

func filter(lines []string, winner, pos int) []string {
	var filtered []string
	for _, line := range lines {
		char := line[pos : pos+1]
		n, err := strconv.Atoi(char)
		if err != nil {
			log.Fatalf("conversion to int failed: %v", err)
		}
		if n == winner {
			filtered = append(filtered, line)
		}
	}
	return filtered
}

func main() {

	inFile, err := os.Open("./day_03/input.dat")
	if err != nil {
		log.Fatalf("loading input data failed: %v", err)
	}
	defer inFile.Close()

	var input1 []string
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()
		input1 = append(input1, line)
	}
	input2 := make([]string, len(input1))
	copy(input2, input1)

	for i := 0; i < 12; i++ {
		if len(input1) > 1 {
			onesWinner := count(input1, i, true)
			input1 = filter(input1, onesWinner, i)
		}

		if len(input2) > 1 {
			zerosWinner := count(input2, i, false)
			input2 = filter(input2, zerosWinner, i)
		}
	}

	if len(input1) != 1 || len(input2) != 1 {
		log.Fatalf("Wrong calculation")
	}

	fmt.Println(input1[0])
	oxygen, err := strconv.ParseInt(input1[0], 2, 64)
	if err != nil {
		log.Fatalf("ParseInt failed: %d %v", input1[0], err)
	}
	fmt.Println(input2[0])
	co2, err := strconv.ParseInt(input2[0], 2, 64)
	if err != nil {
		log.Fatalf("ParseInt failed: %d %v", input2[0], err)
	}

	fmt.Println("Result: ", oxygen*co2)
}
