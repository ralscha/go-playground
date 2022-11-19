package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	inFile, err := os.Open("./day_14/input.dat")
	if err != nil {
		log.Fatalf("loading input data failed: %v", err)
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	scanner.Scan()
	template := scanner.Text()
	rules := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		pos := strings.Index(line, " -> ")
		if pos != -1 {
			rules[line[:pos]] = line[pos+4:]
		}
	}

	for i := 0; i < 10; i++ {
		newTemplate := template[:1]
		ix := 0
		for ix < len(template)-1 {
			k := template[ix : ix+2]
			if v, ok := rules[k]; ok {
				newTemplate = newTemplate + v + k[1:]
			}
			ix++
		}

		template = newTemplate
	}

	counts := make(map[string]int)
	for _, c := range template {
		counts[string(c)]++
	}

	min := ""
	max := ""
	for k, v := range counts {
		if min == "" || v < counts[min] {
			min = k
		}
		if max == "" || v > counts[max] {
			max = k
		}
	}

	fmt.Println("Result: ", counts[max]-counts[min])
}
