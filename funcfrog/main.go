package main

import (
	"fmt"
	"github.com/koss-null/funcfrog/pkg/pipe"
	"github.com/koss-null/funcfrog/pkg/pipies"
	"strconv"
)

func main() {
	a := []int{3, 6, 9, 12, 15, 18, 21, 24, 27, 30}
	res := pipe.Slice(a).
		Map(func(x int) int { return x * x }).
		Map(func(x int) int { return x + 1 }).
		Filter(func(x *int) bool { return *x > 100 }).
		Filter(func(x *int) bool { return *x < 1000 }).
		Parallel(12).
		Do()

	fmt.Println(res)

	p := pipe.Func(func(i int) (int, bool) {
		if i < 10 {
			return i * i, true
		}
		return 0, false
	}).Take(5).Do()
	fmt.Println(p)

	p = pipe.Slice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
		Filter(func(x *int) bool { return *x%2 == 0 }).
		Map(func(x int) int { return len(strconv.Itoa(x)) }).
		Do()
	fmt.Println(p)

	p1 := pipe.Func(func(i int) (float32, bool) {
		return float32(i) * 0.9, true
	}).
		Map(func(x float32) float32 { return x * x }).
		Gen(100).                   // Sort is only availavle on pipes with known length
		Sort(pipies.Less[float32]). // pipe.Less(x, y *T) bool is available to all comparables
		Parallel(12).
		Do()
	fmt.Println(p1)
}
