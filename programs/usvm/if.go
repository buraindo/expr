package main

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func add(i int) int {
	i = i + 2
	if i < 5 {
		return i
	}
	return i + 1
}

func gcd(a, b int64) int64 {
	if a == 0 || b == 0 {
		return -1
	}
	if a == 1 || b == 1 {
		return 1
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {}
