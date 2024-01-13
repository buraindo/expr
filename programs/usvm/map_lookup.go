package main

func lookup(m map[int]int, k int) int {
	delete(m, k)
	v := m[k]
	if v == 123 {
		return -1
	}
	return v
}

func main() {}
