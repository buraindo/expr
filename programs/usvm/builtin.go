package main

func appendSimple(a []int) []int {
	if len(a) == 0 || a == nil {
		return []int{1, 2}
	}

	var c []int = nil
	b := append(c, 5, 6)
	return append(a, b...)
}

func deleteSimple(a map[int]int, k int) int {
	l := len(a)

	if l < 5 {
		panic("too smol map")
	}

	delete(a, k)
	if len(a) == l {
		panic("not found")
	}

	return len(a)
}

func copySimple(a []int) []int {
	if len(a) == 0 {
		return a
	}

	a[0] = 5
	b := make([]int, len(a))
	copy(b, a)
	return b
}

func main() {}
