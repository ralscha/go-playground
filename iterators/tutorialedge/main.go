package main

import (
	"fmt"
	"iter"
	"time"
)

func Countdown(v int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := v; i >= 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

type Employee struct {
	Name   string
	Salary int
}

var Employees = []Employee{
	{Name: "Elliot", Salary: 4},
	{Name: "Donna", Salary: 5},
}

func EmployeeIterator(e []Employee) iter.Seq2[int, Employee] {
	return func(yield func(int, Employee) bool) {
		for i := 0; i <= len(e)-1; i++ {
			if !yield(i, e[i]) {
				return
			}
		}
	}
}

func SleepyIterator[E any](e []E) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i := 0; i <= len(e)-1; i++ {
			time.Sleep(2 * time.Second)
			if !yield(i, e[i]) {
				return
			}
		}
	}
}

func main() {
	Countdown(10)(func(i int) bool {
		fmt.Println(i)
		return true
	})

	for x := range Countdown(10) {
		fmt.Println(x)
	}

	for ix, e := range EmployeeIterator(Employees) {
		fmt.Println(ix, e)
	}

	for i, employee := range SleepyIterator(Employees) {
		fmt.Printf("%d: %+v\n", i, employee)
	}

	for i, val := range SleepyIterator([]int{1, 2, 3, 4}) {
		fmt.Printf("%d: %+v\n", i, val)
	}
}
