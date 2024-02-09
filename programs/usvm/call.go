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
	sort.Ints(data)
	return data
}

func main() {}
