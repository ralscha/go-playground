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

	inFile, err := os.Open("./day_06/input.dat")
	if err != nil {
		log.Fatalf("loading input data failed: %v", err)
	}
	defer func(inFile *os.File) {
		err := inFile.Close()
		if err != nil {
			log.Fatalf("closing input file failed: %v", err)
		}
	}(inFile)

	scanner := bufio.NewScanner(inFile)
	scanner.Scan()
	line := scanner.Text()
	splitted := strings.Split(line, ",")

	var eco [9]int
	for _, s := range splitted {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("conversion failed: %s %v", s, err)
		}

		eco[n] = eco[n] + 1
	}

	for i := 0; i < 256; i++ {
		var neweco [9]int
		for j := 1; j < len(eco); j++ {
			neweco[j-1] = eco[j]
		}
		neweco[6] += eco[0]
		neweco[8] += eco[0]
		eco = neweco
	}

	total := 0
	for _, f := range eco {
		total += f
	}
	fmt.Println("Result: ", total)
}
