package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type StationData struct {
	Name  string
	Min   float64
	Max   float64
	Sum   float64
	Count int
}

func run() {
	started := time.Now()
	data := make(map[string]*StationData)

	file, err := os.Open("../measurements.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ";")
		name := parts[0]
		tempStr := strings.Trim(parts[1], "\n")

		temperature, err := strconv.ParseFloat(tempStr, 64)
		if err != nil {
			panic(err)
		}

		station, ok := data[name]
		if !ok {
			data[name] = &StationData{name, temperature, temperature, temperature, 1}
		} else {
			if temperature < station.Min {
				station.Min = temperature
			}
			if temperature > station.Max {
				station.Max = temperature
			}
			station.Sum += temperature
			station.Count++
		}
	}

	fmt.Printf("%0.6f", time.Since(started).Seconds())

	printResult(data)
}

func printResult(data map[string]*StationData) {
	result := make(map[string]*StationData, len(data))
	keys := make([]string, 0, len(data))
	for _, v := range data {
		keys = append(keys, v.Name)
		result[v.Name] = v
	}
	sort.Strings(keys)

	fmt.Print("{")
	for _, k := range keys {
		v := result[k]
		fmt.Printf("%s=%.1f/%.1f/%.1f, ", k, v.Min, v.Sum/float64(v.Count), v.Max)
	}
	fmt.Print("}\n")
}

func main() {
	run()
}
