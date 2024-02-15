package main

import "sort"

func alloc(l int) []int {
	arr := make([]int, l)
	if l < 5 || l > 10 {
		return arr
	}
	arr[3] = 111
	return arr
}

func overwrite(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	arr[0] = 152
	return arr[0]
}

func sum(nums []int) int {
	if len(nums) < 5 {
		return -228
	}
	res := 0
	for i := range nums {
		if nums[i] > 0 {
			res += nums[i]
		} else if nums[i] < 0 {
			res -= nums[i]
		}
	}
	return res
}

func first(nums []int) int {
	if len(nums) < 2 || len(nums) > 20 {
		return -1
	}
	if nums[0] == 1 && nums[1] == 2 {
		return 1
	}
	return 0
}

func sumArray(nums []int) int {
	if len(nums) < 5 {
		return -228
	}
	res := 0
	for i := range nums {
		if nums[i] > 0 {
			res += nums[i]
		} else if nums[i] < 0 {
			res -= nums[i]
		}
	}
	return res
}

func compare(nums []int) int {
	i := nums[0]
	j := nums[1]
	if i < j {
		return i
	}
	return j
}

func castSlice(arr []int) []int {
	s := sort.IntSlice(arr)
	x := s[0]
	s[0] = x + 2
	return []int(s)
}

func main() {}
