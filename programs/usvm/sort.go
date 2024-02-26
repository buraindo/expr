package main

import "sort"

func count(data []int) []int {
	l := len(data)
	if l == 0 || l > 3 {
		return data
	}

	var aMin, aMax = data[0], data[0]
	for _, x := range data {
		if x < aMin {
			aMin = x
		}
		if x > aMax {
			aMax = x
		}
	}
	if aMax > 1000 || aMax < -1000 || aMin < -1000 || aMin > 1000 {
		return data
	}

	length := aMax - aMin + 1
	if length > 1000 {
		return data
	}
	if l < 3 {
		return data
	}

	allEqual := true
	for _, x := range data {
		if x != data[0] {
			allEqual = false
			break
		}
	}
	if allEqual {
		return data
	}

	count := make([]int, length)
	for _, x := range data {
		count[x-aMin]++
	}

	z := 0
	out := make([]int, len(data))
	for i, c := range count {
		for c > 0 {
			out[z] = i + aMin
			z++
			c--
		}
	}

	return out
}

func bubble(arr []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
	}
	return arr
}

func bubbleCast(a []int) []int {
	arr := sort.IntSlice(a)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < arr.Len()-1; i++ {
			if arr.Less(i+1, i) {
				arr.Swap(i+1, i)
				swapped = true
			}
		}
	}
	return arr
}

func main() {}
