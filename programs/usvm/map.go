package main

func alloc(l int) map[int]int64 {
	m := make(map[int]int64, l)
	m[3] = 111
	m[-226] = 13
	m[0] = -1
	return m
}

func lookup(m map[int]int, k int) int {
	v := m[k]
	if v == 123 {
		return -1
	}
	return v
}

func lookupComma(m map[int]int, k int) int {
	v, ok := m[k]
	if !ok {
		return 0
	}
	if v == 123 {
		return -1
	}
	return v
}

func update(m map[int]int, k, v int) int {
	vOld := m[k]
	m[k] = v
	if v > vOld {
		return v
	}
	return vOld
}

func main() {}
