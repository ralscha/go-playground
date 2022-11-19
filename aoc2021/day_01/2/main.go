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
	lastWindow := -1
	windowCount := 0
	windowSum := 0

	for ix, nums := range numbers {
		num, err := strconv.Atoi(nums)
		if err != nil {
			log.Fatalf("conversion failed: %s %v", nums, err)
		}
		windowSum += num
		if windowCount == 3 {
			if windowSum > lastWindow && lastWindow != -1 {
				increasing += 1
			}
			lastWindow = windowSum

			firstNumOfWindow, err := strconv.Atoi(numbers[ix-2])
			if err != nil {
				log.Fatalf("conversion failed: %s %v", nums, err)
			}
			windowSum = windowSum - firstNumOfWindow
		} else {
			windowCount += 1
			windowSum += num
		}

	}
	fmt.Printf("increasing %d\n", increasing)
}
