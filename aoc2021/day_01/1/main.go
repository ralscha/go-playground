package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./day_01/input.dat")
	if err != nil {
		log.Fatalf("loading input data failed: %v", err)
	}

	numbers := strings.Fields(string(input))
	increasing := 0
	lastNum := -1

	for _, nums := range numbers {
		num, err := strconv.Atoi(nums)
		if err != nil {
			log.Fatalf("conversion failed: %s %v", nums, err)
		}

		if lastNum != -1 && num > lastNum {
			increasing += 1
		}
		lastNum = num
	}
	fmt.Printf("increasing %d\n", increasing)
}
