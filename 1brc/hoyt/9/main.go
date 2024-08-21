package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"sort"
	"time"
)

type r9Stats struct {
	min, max, count int32
	sum             int64
}

func main() {
	start := time.Now()
	err := r9("../measurements.txt", os.Stdout)
	if err != nil {
		fmt.Println(err)
		return
	}
	end := time.Now()
	fmt.Printf("%0.6f\n", end.Sub(start).Seconds())
}

func r9(inputPath string, output io.Writer) error {
	maxGoroutines := runtime.NumCPU()
	parts, err := splitFile(inputPath, maxGoroutines)
	if err != nil {
		return err
	}

	resultsCh := make(chan map[string]*r9Stats)
	for _, part := range parts {
		go r9ProcessPart(inputPath, part.offset, part.size, resultsCh)
	}

	totals := make(map[string]*r9Stats)
	for i := 0; i < len(parts); i++ {
		result := <-resultsCh
		for station, s := range result {
			ts := totals[station]
			if ts == nil {
				totals[station] = s
				continue
			}
			ts.min = min(ts.min, s.min)
			ts.max = max(ts.max, s.max)
			ts.sum += s.sum
			ts.count += s.count
		}
	}

	stations := make([]string, 0, len(totals))
	for station := range totals {
		stations = append(stations, station)
	}
	sort.Strings(stations)

	_, _ = fmt.Fprint(output, "{")
	for i, station := range stations {
		if i > 0 {
			_, _ = fmt.Fprint(output, ", ")
		}
		s := totals[station]
		mean := float64(s.sum) / float64(s.count) / 10
		_, _ = fmt.Fprintf(output, "%s=%.1f/%.1f/%.1f", station, float64(s.min)/10, mean, float64(s.max)/10)
	}
	_, _ = fmt.Fprint(output, "}\n")
	return nil
}

func r9ProcessPart(inputPath string, fileOffset, fileSize int64, resultsCh chan map[string]*r9Stats) {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.Seek(fileOffset, io.SeekStart)
	if err != nil {
		panic(err)
	}
	f := io.LimitedReader{R: file, N: fileSize}

	type item struct {
		key  []byte
		stat *r9Stats
	}
	const numBuckets = 1 << 17        // number of hash buckets (power of 2)
	items := make([]item, numBuckets) // hash buckets, linearly probed
	size := 0                         // number of active items in items slice

	buf := make([]byte, 1024*1024)
	readStart := 0
	for {
		n, err := f.Read(buf[readStart:])
		if err != nil && err != io.EOF {
			panic(err)
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
			const (
				// FNV-1 64-bit constants from hash/fnv.
				offset64 = 14695981039346656037
				prime64  = 1099511628211
			)

			var station, after []byte
			hash := uint64(offset64)
			i := 0
			for ; i < len(chunk); i++ {
				c := chunk[i]
				if c == ';' {
					station = chunk[:i]
					after = chunk[i+1:]
					break
				}
				hash ^= uint64(c) // FNV-1a is XOR then *
				hash *= prime64
			}
			if i == len(chunk) {
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

			hashIndex := int(hash & (numBuckets - 1))
			for {
				if items[hashIndex].key == nil {
					// Found empty slot, add new item (copying key).
					key := make([]byte, len(station))
					copy(key, station)
					items[hashIndex] = item{
						key: key,
						stat: &r9Stats{
							min:   temp,
							max:   temp,
							sum:   int64(temp),
							count: 1,
						},
					}
					size++
					if size > numBuckets/2 {
						panic("too many items in hash table")
					}
					break
				}
				if bytes.Equal(items[hashIndex].key, station) {
					// Found matching slot, add to existing stats.
					s := items[hashIndex].stat
					s.min = min(s.min, temp)
					s.max = max(s.max, temp)
					s.sum += int64(temp)
					s.count++
					break
				}
				// Slot already holds another key, try next slot (linear probe).
				hashIndex++
				if hashIndex >= numBuckets {
					hashIndex = 0
				}
			}
		}

		readStart = copy(buf, remaining)
	}

	result := make(map[string]*r9Stats, size)
	for _, item := range items {
		if item.key == nil {
			continue
		}
		result[string(item.key)] = item.stat
	}
	resultsCh <- result
}

type part struct {
	offset, size int64
}

func splitFile(inputPath string, numParts int) ([]part, error) {
	const maxLineLength = 100

	f, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	st, err := f.Stat()
	if err != nil {
		return nil, err
	}
	size := st.Size()
	splitSize := size / int64(numParts)

	buf := make([]byte, maxLineLength)

	parts := make([]part, 0, numParts)
	offset := int64(0)
	for i := 0; i < numParts; i++ {
		if i == numParts-1 {
			if offset < size {
				parts = append(parts, part{offset, size - offset})
			}
			break
		}

		seekOffset := max(offset+splitSize-maxLineLength, 0)
		_, err := f.Seek(seekOffset, io.SeekStart)
		if err != nil {
			return nil, err
		}
		n, _ := io.ReadFull(f, buf)
		chunk := buf[:n]
		newline := bytes.LastIndexByte(chunk, '\n')
		if newline < 0 {
			return nil, fmt.Errorf("newline not found at offset %d", offset+splitSize-maxLineLength)
		}
		remaining := len(chunk) - newline - 1
		nextOffset := seekOffset + int64(len(chunk)) - int64(remaining)
		parts = append(parts, part{offset, nextOffset - offset})
		offset = nextOffset
	}
	return parts, nil
}
