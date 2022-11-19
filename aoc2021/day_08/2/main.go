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

	zero := [7]bool{true, true, true, false, true, true, true}
	one := [7]bool{false, false, true, false, false, true, false}
	two := [7]bool{true, false, true, true, true, false, true}
	three := [7]bool{true, false, true, true, false, true, true}
	four := [7]bool{false, true, true, true, false, true, false}
	five := [7]bool{true, true, false, true, false, true, true}
	six := [7]bool{true, true, false, true, true, true, true}
	seven := [7]bool{true, false, true, false, false, true, false}
	eight := [7]bool{true, true, true, true, true, true, true}
	nine := [7]bool{true, true, true, true, false, true, true}

	numbers := [10][7]bool{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}
	allAgPermutations := agPermutations()

	sum := 0
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()
		pos := strings.IndexAny(line, "|")
		ins := strings.Fields(line[:pos])
		for _, agPerm := range allAgPermutations {
			var ok [10]bool
			for _, in := range ins {
				num := convertToSegment(agPerm, in)
				number := findPattern(numbers, num)
				if number == -1 || ok[number] {
					break
				} else {
					ok[number] = true
				}
			}

			if allTrue(ok) {
				pos := strings.IndexAny(line, "|")
				outs := strings.Fields(line[pos+1:])
				outnumbers := ""
				for _, out := range outs {
					num := convertToSegment(agPerm, out)
					outnum := findPattern(numbers, num)
					outnumbers += strconv.Itoa(outnum)
				}
				outnumbersi, err := strconv.Atoi(outnumbers)
				if err != nil {
					log.Panicln(err)
				}
				sum += outnumbersi
				break
			}
		}
	}

	fmt.Println("Result: ", sum)
}

func convertToSegment(agPerm []string, str string) [7]bool {
	var num [7]bool
	for _, c := range str {
		pos := index(agPerm, string(c))
		num[pos] = true
	}
	return num
}

func allTrue(ok [10]bool) bool {
	for _, o := range ok {
		if !o {
			return false
		}
	}
	return true
}

func findPattern(numbers [10][7]bool, num [7]bool) int {
	for numberix, number := range numbers {
		allOk := true
		for i := 0; i < len(number); i++ {
			if number[i] != num[i] {
				allOk = false
				break
			}
		}
		if allOk {
			return numberix
		}
	}
	return -1
}

func index(arr []string, p string) int {
	for ix, v := range arr {
		if v == p {
			return ix
		}
	}
	return -1
}

func agPermutations() [][]string {
	arr := []string{"a", "b", "c", "d", "e", "f", "g"}

	var helper func([]string, int)
	var res [][]string

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
