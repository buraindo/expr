package main

func first(nums []int) int {
	if len(nums) < 2 || len(nums) > 20 {
		return -1
	}
	if nums[0] == 1 && nums[1] == 2 {
		return 1
	}
	return 0
}

func main() {}
