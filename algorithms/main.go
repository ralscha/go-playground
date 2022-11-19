package main

import "fmt"

func main() {
	s := "abc"
	perm := permutation(s)
	fmt.Println(perm)

	permIter := permutationIterative(s)
	fmt.Println(permIter)

	// solveEightQueen()

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(binarySearch(arr, 5))
}
