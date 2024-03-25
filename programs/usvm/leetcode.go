package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i := range nums {
		n := nums[i]
		m[target-n] = i
	}
	for i := range nums {
		n := nums[i]
		if _, ok := m[n]; !ok {
			continue
		}
		if m[n] != i {
			return []int{i, m[n]}
		}
	}

	return nil
}

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, k)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		if m[nums[i]] >= 2 {
			return true
		}
		if i-k >= 0 {
			m[nums[i-k]]--
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {}
