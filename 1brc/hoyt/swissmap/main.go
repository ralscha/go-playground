package main

import (
	"bufio"
	"fmt"
	"github.com/dolthub/swiss"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type stats struct {
	min, max, sum float64
	count         int64
}

func main() {
	measurementsPath := "../measurements.txt"

	start := time.Now()
	stationStats, stations, err := r(measurementsPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	end := time.Now()
	fmt.Printf("%0.6f\n\n", end.Sub(start).Seconds())

	output := os.Stdout
	_, _ = fmt.Fprint(output, "{")
	for i, station := range stations {
		if i > 0 {
			_, _ = fmt.Fprint(output, ", ")
		}
		s, _ := stationStats.Get(station)
		mean := s.sum / float64(s.count)
		_, _ = fmt.Fprintf(output, "%s=%.1f/%.1f/%.1f", station, s.min, mean, s.max)
	}
	_, _ = fmt.Fprint(output, "}\n")

}

func r(inputPath string) (*swiss.Map[string, *stats], []string, error) {

	f, err := os.Open(inputPath)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	stationStats := swiss.NewMap[string, *stats](42)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		station, tempStr, hasSemi := strings.Cut(line, ";")
		if !hasSemi {
			continue
		}

		temp, err := strconv.ParseFloat(tempStr, 64)
		if err != nil {
			return nil, nil, err
		}

		s, exists := stationStats.Get(station)
		if !exists {
			stationStats.Put(station, &stats{
				min:   temp,
				max:   temp,
				sum:   temp,
				count: 1,
			})
		} else {
			s.min = min(s.min, temp)
			s.max = max(s.max, temp)
			s.sum += temp
			s.count++
		}
	}

	stations := make([]string, 0, stationStats.Count())
	stationStats.Iter(func(station string, _ *stats) bool {
		stations = append(stations, station)
		return false
	})

	sort.Strings(stations)

	return stationStats, stations, nil
}
