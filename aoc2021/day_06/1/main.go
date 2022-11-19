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

	var eco []*byte
	for _, s := range splitted {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("conversion failed: %s %v", s, err)
		}
		b := byte(n)
		eco = append(eco, &b)
	}

	for i := 0; i < 80; i++ {
		for _, f := range eco {
			if *f > 0 {
				*f = *f - 1
			} else {
				*f = 6
				b := byte(8)
				eco = append(eco, &b)
			}
		}
	}

	fmt.Println("Result: ", len(eco))
}
