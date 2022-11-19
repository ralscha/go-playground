package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	inFile, err := os.Open("./day_08/input.dat")
	if err != nil {
		log.Fatalf("loading input data failed: %v", err)
	}
	defer func(inFile *os.File) {
		err := inFile.Close()
		if err != nil {
			log.Fatalf("closing input file failed: %v", err)
		}
	}(inFile)

	count := 0
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()

		pos := strings.IndexAny(line, "|")
		ins := strings.Fields(line[pos+1:])
		for _, in := range ins {
			l := len(in)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				count++
			}
		}
	}

	fmt.Println("Result: ", count)
}
