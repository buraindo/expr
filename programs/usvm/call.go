package main

import (
	"sort"
)

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min3(a, b, c int) int {
	return min2(min2(a, b), c)
}

func stdSort(data []int) []int {
	if len(data) != 7 {
		return data
	}

	sort.Ints(data)
	return data
}

func swap(data sort.Interface, a, b int) (int, int) {
	if data.Less(b, a) {
		return b, a
	}
	return a, b
}

func callSwap(arr []int) (int, int) {
	return swap(sort.IntSlice(arr), 1, 3)
}

func callAnonClosure(i int) int {
	if i == 0 {
		return -1
	}
	f := func(j int) int {
		return i + j
	}
	return f(2)
}

type Person struct {
	Age int
}

func (p Person) getAge() int {
	return p.Age
}

func callBoundClosure(i int) int {
	p := Person{Age: i + 2}
	f := p.getAge
	return f()
}

func main() {}
