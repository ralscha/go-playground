package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Alice", "Bob", "Vera", "Zac", "Zac2"}

	aIndex := slices.Index(names, "Alice")

	names = slices.Replace(names, aIndex, aIndex+1)
	fmt.Printf("Value %v\n", names)

	names = slices.Replace(names, aIndex, aIndex+1, "John")
	fmt.Printf("Value %v\n", names)

	names = slices.Replace(names, aIndex, aIndex+1, "Mary", "Jane")
	fmt.Printf("Value %v\n", names)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slices.Index(numbers, 5))
	fmt.Println(slices.Index(numbers, 10))
	fmt.Println(slices.Max(numbers))
	fmt.Println(slices.Min(numbers))
	fmt.Println(slices.Contains(numbers, 5))
	fmt.Println(slices.Contains(numbers, 10))
	fmt.Println(slices.Compare(numbers, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Println(slices.Compare(numbers, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))

	fmt.Println(slices.Compact(numbers))
	fmt.Println(slices.Compact([]int{1, 1, 1, 1, 2, 2, 2, 3, 4, 5, 6, 2, 7, 7, 8, 9, 0}))

	fmt.Println(slices.BinarySearch(numbers, 5))

	numbers = []int{1}
	numbers = append(numbers, 2)
	numbers = append(numbers, 3)
	fmt.Println("before")
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
	clipped := slices.Clip(numbers)
	fmt.Println("after clip")
	fmt.Println(len(clipped))
	fmt.Println(cap(clipped))
	grown := slices.Grow(numbers, 3)
	fmt.Println("after grow")
	fmt.Println(len(grown))
	fmt.Println(cap(grown))

	numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slices.Delete(numbers, 2, 5))
	fmt.Println(slices.Equal(numbers, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}))

	numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	numbers = slices.Insert(numbers, 2, 5, 10, 11, 12)
	fmt.Println(numbers)
	fmt.Println(slices.IsSorted(numbers))
	slices.Sort(numbers)
	fmt.Println(slices.IsSorted(numbers))
	fmt.Println(numbers)
}
