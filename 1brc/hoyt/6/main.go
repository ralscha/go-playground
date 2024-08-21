package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"time"
)

type stats struct {
	min, max, count int32
	sum             int64
}

func main() {
	measurementsPath := "../measurements.txt"

	start := time.Now()
	stationStats, stations, err := r6(measurementsPath)
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
		s := stationStats[station]
		mean := float64(s.sum) / float64(s.count) / 10
		_, _ = fmt.Fprintf(output, "%s=%.1f/%.1f/%.1f",
			station, float64(s.min)/10, mean, float64(s.max)/10)
	}
	_, _ = fmt.Fprint(output, "}\n")

}

func r6(inputPath string) (map[string]*stats, []string, error) {

	f, err := os.Open(inputPath)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	stationStats := make(map[string]*stats)

	buf := make([]byte, 1024*1024)
	readStart := 0
	for {
		n, err := f.Read(buf[readStart:])
		if err != nil && err != io.EOF {
			return nil, nil, err
		}
		if readStart+n == 0 {
			break
		}
		chunk := buf[:readStart+n]

		newline := bytes.LastIndexByte(chunk, '\n')
		if newline < 0 {
			break
		}
		remaining := chunk[newline+1:]
		chunk = chunk[:newline+1]

		for {
			station, after, hasSemi := bytes.Cut(chunk, []byte(";"))
			if !hasSemi {
				break
			}

			index := 0
			negative := false
			if after[index] == '-' {
				negative = true
				index++
			}
			temp := int32(after[index] - '0')
			index++
			if after[index] != '.' {
				temp = temp*10 + int32(after[index]-'0')
				index++
			}
			index++ // skip '.'
			temp = temp*10 + int32(after[index]-'0')
			index += 2 // skip last digit and '\n'
			if negative {
				temp = -temp
			}
			chunk = after[index:]

			s := stationStats[string(station)]
			if s == nil {
				stationStats[string(station)] = &stats{
					min:   temp,
					max:   temp,
					sum:   int64(temp),
					count: 1,
				}
			} else {
				s.min = min(s.min, temp)
				s.max = max(s.max, temp)
				s.sum += int64(temp)
				s.count++
			}
		}

		readStart = copy(buf, remaining)
	}

	stations := make([]string, 0, len(stationStats))
	for station := range stationStats {
		stations = append(stations, station)
	}
	sort.Strings(stations)

	return stationStats, stations, nil
}