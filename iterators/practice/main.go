package main

import (
	"fmt"
	"iter"
	"slices"
)

func main() {

	seq := []int{1, 2, 3, 4, 5}
	for item := range FirstN(slices.Values(seq), 3) {
		fmt.Print(item)
		fmt.Print(" ")
	}
	fmt.Println()

	for item := range SkipFirstN(slices.Values(seq), 2) {
		fmt.Print(item)
		fmt.Print(" ")
	}
	fmt.Println()

	for item := range SkipFirstN(FirstN(slices.Values(seq), 3), 1) {
		fmt.Print(item)
		fmt.Print(" ")
	}
	fmt.Println()

	fmt.Println("Chunk:")
	for item := range Chunk(slices.Values(seq), 2) {
		for subItem := range item {
			fmt.Print(subItem)
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("Chunk2:")
	for item := range Chunk2(slices.All(seq), 2) {
		for ix, item := range item {
			fmt.Print(ix)
			fmt.Print(": ")
			fmt.Print(item)
			fmt.Print(" ")
		}
		fmt.Println()
	}

	fmt.Println("SeqToSeq2:")
	for k, v := range SeqToSeq2(slices.Values(seq), func(v int) string {
		return fmt.Sprintf("key-%d", v)
	}) {
		fmt.Print(k)
		fmt.Print(":")
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println()

	fmt.Println("Map:")
	for item := range Map(slices.Values(seq), func(t int) int {
		return t * 2
	}) {
		fmt.Print(item)
		fmt.Print(" ")
	}
	fmt.Println()

	fmt.Println("Map2:")
	for k, v := range Map2(slices.All(seq), func(k int) string {
		return fmt.Sprintf("key-%d", k)
	}, func(v int) string {
		return fmt.Sprintf("value-%d", v)
	}) {
		fmt.Print(k)
		fmt.Print(":")
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println()
}

// FirstN takes an iter.Seq[T] and returns a new iter.Seq[T] that
// yields only the first 'limit' items without creating intermediate slices.
func FirstN[T any](original iter.Seq[T], limit int) iter.Seq[T] {
	if limit <= 0 {
		return func(yield func(T) bool) {}
	}

	return func(yield func(T) bool) {
		next, stop := iter.Pull[T](original)
		defer stop()

		for range limit {
			val, ok := next()
			if !ok {
				break
			}
			if !yield(val) {
				return
			}
		}
	}
}

func SkipFirstN[T any](seq iter.Seq[T], skip int) iter.Seq[T] {
	if skip <= 0 {
		return seq
	}

	return func(yield func(T) bool) {
		next, stop := iter.Pull[T](seq)
		defer stop()

		for i := 0; i < skip; i++ {
			_, ok := next()
			if !ok {
				break
			}
		}
		for {
			v, ok := next()
			if !ok {
				break
			}
			if !yield(v) {
				return
			}
		}
	}
}

// Chunk returns an iterator over consecutive sub-slices of up to n elements of s.
// All but the last iter.Seq chunk will have size n.
// Chunk panics if n is less than 1.
func Chunk[T any](sq iter.Seq[T], size int) iter.Seq[iter.Seq[T]] {
	if size < 1 {
		panic("size must be >= 1")
	}

	return func(yield func(s iter.Seq[T]) bool) {
		next, stop := iter.Pull[T](sq)
		defer stop()

		for {
			firstVal, firstOk := next()
			if !firstOk {
				// No more items in the sequence
				break
			}

			chunk := make([]T, 0, size)
			chunk = append(chunk, firstVal)

			for range size - 1 {
				val, ok := next()
				if !ok {
					break
				}
				chunk = append(chunk, val)
			}

			if !yield(slices.Values(chunk)) {
				return
			}
		}
	}
}

// Chunk2 returns an iterator over consecutive sub-slices of up to n elements of s.
// All but the last iter.Seq chunk will have size n.
// Chunk2 panics if n is less than 1.
func Chunk2[K any, V any](sq iter.Seq2[K, V], size int) iter.Seq[iter.Seq2[K, V]] {
	if size < 1 {
		panic("size must be >= 1")
	}

	return func(yield func(s iter.Seq2[K, V]) bool) {
		next, stop := iter.Pull2[K, V](sq)
		defer stop()

		for {
			// Get the first item for this chunk
			firstK, firstV, firstOk := next()
			if !firstOk {
				// No more items in the sequence
				break
			}

			// Create a chunk with the first item and up to size-1 more items
			keys := make([]K, 0, size)
			values := make([]V, 0, size)

			keys = append(keys, firstK)
			values = append(values, firstV)

			// Add more items to the chunk up to the size limit
			for range size - 1 {
				k, v, ok := next()
				if !ok {
					// No more items in the sequence
					break
				}
				keys = append(keys, k)
				values = append(values, v)
			}

			// Create a sequence from the chunk
			chunkSeq := func(yield func(K, V) bool) {
				for i := 0; i < len(keys); i++ {
					if !yield(keys[i], values[i]) {
						return
					}
				}
			}

			// Yield the chunk sequence
			if !yield(chunkSeq) {
				return
			}
		}
	}
}

func SeqToSeq2[K any, V any](is iter.Seq[V], keyFunc func(v V) K) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		next, stop := iter.Pull[V](is)
		defer stop()

		for {
			v, ok := next()
			if !ok {
				break
			}
			k := keyFunc(v)
			if !yield(k, v) {
				return
			}
		}
	}
}

func Map[T any, R any](it iter.Seq[T], mapFunc func(t T) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		for i := range it {
			if !yield(mapFunc(i)) {
				return
			}
		}
	}
}

func Map2[K any, V any, KR any, VR any](it iter.Seq2[K, V], keyFunc func(k K) KR, valFunc func(v V) VR) iter.Seq2[KR, VR] {
	return func(yield func(KR, VR) bool) {
		for k, v := range it {
			if !yield(keyFunc(k), valFunc(v)) {
				return
			}
		}
	}
}
