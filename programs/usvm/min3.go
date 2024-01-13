package main

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min3(a, b, c int) int {
	return min2(min2(a, b), c)
}

// func min3(a, b, c int) int {
// 	if a <= b && a <= c {
// 		return a
// 	}
// 	if b <= c && b <= a {
// 		return b
// 	}
// 	return c
// }

func main() {
	println(min3(2, 1, -1))
}
